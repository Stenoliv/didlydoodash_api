package kanban

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/datatypes"
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
		Type:    utils.KanbanError,
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

	// TODO: Check permissions

	switch input.Type {
	// Kanban
	case utils.EditKanban:
		var editKanban EditKanban
		if err := json.Unmarshal(input.Payload, &editKanban); err != nil {
			c.SendErrorMessage("User error! Wrong input")
			return
		}
		// Get kanban
		kanban, err := daos.GetKanban(editKanban.ID)
		if err != nil || kanban == nil {
			c.SendErrorMessage("User error! Wrong input")
			return
		}
		// Start transaction and update
		tx := db.DB.Begin()
		if err := tx.Model(&kanban).Updates(editKanban.Updates).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to update kanban")
			return
		}
		// turn into response
		payload, err := editKanban.ToJSON()
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to update kanban")
			return
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to update kanban")
			return
		}
		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.EditKanban,
			RoomID:  c.RoomID,
			Payload: payload,
		}
	// Kanban categories
	case utils.NewKanbanCategory: // Create new category functionality
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
	case utils.EditKanbanCategory: // Edit category functionality
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
		// Create raw json of response
		response := &CategoryResponse{
			Category: *category,
		}
		payload, err := response.ToJSON()
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Failed to update kanban category")
			return
		}
		// Try to commit to database
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
	case utils.DeleteKanbanCategory: // Delete cateogry functionality
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
		// Create raw json of response
		deleteCategoryResponse := &CategoryResponse{
			Category: *category,
		}
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
	// Kanban items
	case utils.NewKanbanItem:
		var newItem NewItem
		if err := json.Unmarshal(input.Payload, &newItem); err != nil {
			c.SendErrorMessage("User error! Failed to create new item")
			return
		}
		// create new item variable
		item := &models.KanbanItem{
			KanbanCategoryID: newItem.CategoryID,
			Priority:         datatypes.NONE,
			Title:            newItem.Name,
		}
		// Start transaction and create in database
		tx := db.DB.Begin()
		if err := item.SaveItem(tx); err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to create new item")
			return
		}
		// Create response payload and jsonify
		payload := &ItemResponse{
			Item: *item,
		}
		raw, err := payload.ToJSON()
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to create new item")
			return
		}
		// Try to save transaction on database
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to create new item")
			return
		}
		// Send to hubs broadcast channel
		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.NewKanbanItem,
			RoomID:  c.RoomID,
			Payload: raw,
		}
	case utils.MoveKanbanItem:
		var moveItem MoveItem
		if err := json.Unmarshal(input.Payload, &moveItem); err != nil {
			c.SendErrorMessage("User error! Failed to move item")
			return
		}
		// Check if new category exists
		category, err := daos.GetCategory(moveItem.NewCategoryID)
		if err != nil || category == nil {
			c.SendErrorMessage("User error! Category target doesn't exist")
			return
		}
		item, err := daos.GetKanbanItem(moveItem.ItemID)
		if err != nil || item == nil {
			c.SendErrorMessage("User error! Item to move doesn't exist")
			return
		}
		// Start transaction to update item
		tx := db.DB.Begin()
		if err := tx.Model(&item).Update("kanban_category_id", category.ID).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to move item")
			return
		}
		// Convert to raw json payload
		payload, err := moveItem.ToJSON()
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to move item")
			return
		}
		// Try to save change to database
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to move item")
			return
		}
		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.MoveKanbanItem,
			RoomID:  c.RoomID,
			Payload: payload,
		}
	case utils.EditKanbanItem:
		var updateItem EditItem
		if err := json.Unmarshal(input.Payload, &updateItem); err != nil {
			c.SendErrorMessage("User error! Failed to update item")
			return
		}
		// Check if item exists in database
		item, err := daos.GetKanbanItem(updateItem.ItemID)
		if err != nil || item == nil {
			c.SendErrorMessage("User error! Item not found")
			return
		}
		// Try to delete item
		tx := db.DB.Begin()
		if updateItem.Updates["estimated_time"] == "" {
			updateItem.Updates["estimated_time"] = nil
		}
		if updateItem.Updates["due_date"] == "" {
			updateItem.Updates["due_date"] = nil
		}
		if err := tx.Model(&item).Updates(updateItem.Updates).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to update item")
			return
		}

		// Explicitly reload the updated item from the database
		if err := tx.Model(&item).First(&item).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to retrieve updated item")
			return
		}
		// Try to make response
		response := &ItemResponse{
			Item: *item,
		}
		payload, err := json.Marshal(&response)
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to update item")
			return
		}
		// Try to save transaction
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to update item")
			return
		}
		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.EditKanbanItem,
			RoomID:  c.RoomID,
			Payload: payload,
		}
	case utils.DeleteKanbanItem:
		var deleteItem DeleteItem
		if err := json.Unmarshal(input.Payload, &deleteItem); err != nil {
			c.SendErrorMessage("User error! Failed to remove item")
			return
		}
		// Check if item exists in database
		item, err := daos.GetKanbanItem(deleteItem.ItemID)
		if err != nil || item == nil {
			c.SendErrorMessage("User error! Item not found")
			return
		}
		// Try to delete item
		tx := db.DB.Begin()
		if err := tx.Delete(&item).Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to remove item")
			return
		}
		// Try to make response
		response := &ItemResponse{
			Item: *item,
		}
		payload, err := json.Marshal(&response)
		if err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to remove item")
			return
		}
		// Try to save transaction
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.SendErrorMessage("Server error! Failed to remove item")
			return
		}
		handler.Hub.Broadcast <- &ws.WSMessage{
			Type:    utils.DeleteKanbanItem,
			RoomID:  c.RoomID,
			Payload: payload,
		}
	// Default
	default:
		return
	}
}
