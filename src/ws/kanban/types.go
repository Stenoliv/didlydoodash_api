package kanban

import (
	"DidlyDoodash-api/src/db/models"
	"encoding/json"
)

// Join kanban room message
type JoinMessage struct {
	Kanban  models.Kanban              `json:"kanban"`
	Archive []models.KanbanArchiveItem `json:"archive"`
}

func (m *JoinMessage) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

type EditKanban struct {
	ID       string                 `json:"id" binding:"required"`
	Updates  map[string]interface{} `json:"updates"`
	SenderID string                 `json:"userId"`
}

func (m *EditKanban) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

/**
 * Category
 */

// New category input message
type NewCategory struct {
	Name string `json:"name"`
}

type RestoreKanbanCategory struct {
	ID string `json:"id" binding:"required"`
}

// Edit category websocket message
type EditCategory struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name"`
}

// Delete category websocket message
type DeleteCategory struct {
	ID string `json:"id" binding:"required"`
}

// Category response
type CategoryResponse struct {
	Category models.KanbanCategory `json:"category"`
	SenderID string                `json:"userId"`
}

func (m *CategoryResponse) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

/**
 * Items
 */

// Items
type NewItem struct {
	CategoryID string `json:"categoryId" binding:"required"`
	Name       string `json:"name"`
}

type RestoreKanbanItem struct {
	ItemID string `json:"itemId" binding:"required"`
}

type MoveItem struct {
	OldCategoryID string `json:"oldCategoryId" binding:"required"`
	NewCategoryID string `json:"newCategoryId" binding:"required"`
	ItemID        string `json:"itemId" binding:"required"`
}

func (m *MoveItem) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

type EditItem struct {
	CategoryID string                 `json:"categoryId" binding:"required"`
	ItemID     string                 `json:"itemId" binding:"requied"`
	Updates    map[string]interface{} `json:"updates"`
}

type DeleteItem struct {
	ItemID string `json:"itemId" binding:"required"`
}

type ItemResponse struct {
	Item     models.KanbanItem `json:"item"`
	SenderID string            `json:"userId"`
}

func (m *ItemResponse) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}
