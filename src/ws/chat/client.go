package chat

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/ws"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *ws.WSMessage
	RoomID  string `json:"roomId"`
	UserID  string `json:"userId"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			break
		}

		if err := c.Conn.WriteJSON(message); err != nil {
			break
		}
	}
}

func (c *Client) readMessage(handler *ChatHandler) {
	defer func() {
		handler.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("unexpected error:", err)
			}
			break
		}
		go c.HandleMessage(msg, handler)
	}
}

func (c *Client) HandleMessage(msg []byte, handler *ChatHandler) {
	var input ws.WSMessage
	if err := json.Unmarshal(msg, &input); err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return
	}
	switch input.Type {
	case utils.MessageSend:
		// Try to unmarshal the ws message payload to a valid message struct
		var message MessageSend
		if err := json.Unmarshal(input.Payload, &message); err != nil {
			// Failed to parse input payload
			Err := &ws.WSError{
				Message: "Invalid message input",
			}

			c.Message <- &ws.WSMessage{
				Type:    utils.MessageError,
				RoomID:  input.RoomID,
				Payload: Err.ToJSON(),
			}
			return
		}

		// Start a transaction and save message to database
		tx := db.DB.Begin()
		dbMessage := &models.ChatMessage{
			RoomID:  input.RoomID,
			UserID:  message.ID,
			Message: message.Message,
		}
		if err := dbMessage.SaveMessage(tx); err != nil {
			// Failed to save message to database
			Err := &ws.WSError{
				Message: "Failed to save message",
			}

			c.Message <- &ws.WSMessage{
				Type:    utils.MessageError,
				RoomID:  input.RoomID,
				Payload: Err.ToJSON(),
			}

			tx.Rollback()
			return
		}

		// Accept changes to database
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return
		}

		// Create a response WSMessage
		response := ws.WSMessage{
			Type:    utils.MessageSend,
			RoomID:  input.RoomID,
			Payload: dbMessage.ToJSON(),
		}

		// Broadcast message to the clients in room
		handler.Hub.Broadcast <- &response

		msg := &NewMessage{
			ChatID:  c.RoomID,
			UserID:  c.UserID,
			Message: message.Message,
		}
		notification, err := msg.ToJSON()
		if err != nil {
			return
		}
		handler.NotificationHub.NewNotification(MessageNew, notification)
	case utils.MessageRead:
		var readMessage MessageRead
		if err := json.Unmarshal(input.Payload, &readMessage); err != nil {
			Err := &ws.WSError{
				Message: "Invalid message input",
			}
			c.Message <- &ws.WSMessage{
				Type:    utils.MessageError,
				RoomID:  input.RoomID,
				Payload: Err.ToJSON(),
			}
			return
		}

		if err := updateUserLastReadMessage(c.RoomID, c.UserID, readMessage.MessageID); err != nil {
			Err := &ws.WSError{
				Message: "Failed to update read message",
			}
			c.Message <- &ws.WSMessage{
				Type:    utils.MessageError,
				RoomID:  input.RoomID,
				Payload: Err.ToJSON(),
			}
			return
		}
	case utils.MessageTyping:
		handler.Hub.Broadcast <- &input
	default:
		handler.Hub.Broadcast <- &input
	}
}

// Function to update the user's last read message
func updateUserLastReadMessage(roomID, userID, messageID string) error {
	var member models.ChatMember

	// Fetch the member record from the database
	if err := db.DB.Model(&models.ChatMember{}).
		Where("room_id = ? AND user_id = ?", roomID, userID).
		First(&member).Error; err != nil {
		return err
	}

	// Update the LastMessageID with the new message that the user has read
	member.LastMessageID = &messageID

	// Save the updated member back to the database
	return db.DB.Save(&member).Error
}
