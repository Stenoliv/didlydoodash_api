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

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
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
		go c.HandleMessage(msg, hub)
	}
}

func (c *Client) HandleMessage(msg []byte, hub *Hub) {
	var input ws.WSMessage
	if err := json.Unmarshal(msg, &input); err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return
	}
	/**
	 * Switch and handle different message types
	 */
	switch input.Type {
	case utils.MessageSend:

		// Try to unmarshal the ws message payload to a valid message struct
		var message MessageStruct
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
		tx.Commit()

		// Create a response WSMessage
		response := ws.WSMessage{
			Type:    utils.MessageSend,
			RoomID:  input.RoomID,
			Payload: dbMessage.ToJSON(),
		}

		// Broadcast message to the clients in room
		hub.Broadcast <- &response
	case utils.MessageTyping:
		hub.Broadcast <- &input
	default:
		hub.Broadcast <- &input
	}
}
