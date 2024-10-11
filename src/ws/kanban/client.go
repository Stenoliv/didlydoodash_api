package kanban

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
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

// Function to send error message to user
func (c *Client) SendErrorMessage(title string) {
	Err := &ws.WSError{
		Message: title,
	}
	c.Message <- &ws.WSMessage{
		Type:    utils.MessageError,
		RoomID:  c.RoomID,
		Payload: Err.ToJSON(),
	}
}

// Function that listens for messages from the client
func (c *Client) readMessage(handler *Handler) {
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
		go c.handleMessage(msg, handler)
	}
}

// Function that writes messages to the clients connection
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

// Function to handle messages
func (c *Client) handleMessage(msg []byte, handler *Handler) {
	var input ws.WSMessage
	if err := json.Unmarshal(msg, &input); err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return
	}
	switch input.Type {
	case utils.JoinKanban:
		break
	case utils.NewKanbanCategory:
		var newCategoryData NewCategory
		if err := json.Unmarshal(input.Payload, &newCategoryData); err != nil {
			c.SendErrorMessage("Failed to create new category")
			return
		}

		category := &models.KanbanCategory{
			KanbanID: c.RoomID,
			Name:     newCategoryData.Name,
		}

		tx := db.DB.Begin()
		if err := category.SaveCategory(tx); err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to create new category")
			return
		}

		payload := &CategoryResponse{
			Category: *category,
		}

		raw, err := payload.ToJSON()
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to create new category")
			return
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to create new category")
			return
		}

		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.NewKanbanCategory,
			RoomID:  c.RoomID,
			Payload: raw,
		}
	case utils.EditKanbanCategory:
		var editInput EditCategory
		if err := json.Unmarshal(input.Payload, &editInput); err != nil {
			c.SendErrorMessage("Failed to update kanban category")
			return
		}

		category, err := daos.GetCategory(editInput.ID)
		if err != nil {
			c.SendErrorMessage("Failed to update kanban category")
			return
		}

		// Start transaction to database
		tx := db.DB.Begin()
		if err := tx.Model(&category).Update("name", editInput.Name).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to update kanban category")
			return
		}

		response := &CategoryResponse{
			Category: *category,
		}

		payload, err := response.ToJSON()
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to update kanban category")
			return
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to update kanban category")
			return
		}

		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.EditKanbanCategory,
			RoomID:  c.RoomID,
			Payload: payload,
		}
	case utils.DeleteKanbanCategory:
		var deleteInput DeleteCategory
		if err := json.Unmarshal(input.Payload, &deleteInput); err != nil {
			c.SendErrorMessage("Failed to delete kanban category")
			return
		}

		// Get category from database
		category, err := daos.GetCategory(deleteInput.ID)
		if err != nil {
			c.SendErrorMessage("Category doesn't exist")
			return
		}

		// Try to delete category
		tx := db.DB.Begin()
		if err := tx.Delete(&category).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to delete kanban category")
			return
		}

		deleteCategoryResponse := &CategoryResponse{
			Category: *category,
		}

		// Create raw json of response
		payload, err := json.Marshal(&deleteCategoryResponse)
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to delete kanban category")
			return
		}

		// Commit changes to database
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to delete kanban category")
			return
		}

		// Send message to clients
		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.DeleteKanbanCategory,
			RoomID:  c.RoomID,
			Payload: payload,
		}
	default:
		return
	}
}
