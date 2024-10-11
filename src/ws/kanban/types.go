package kanban

import (
	"DidlyDoodash-api/src/db/models"
	"encoding/json"
)

// Join kanban room message
type JoinMessage struct {
	Kanban models.Kanban `json:"kanban"`
}

func (m *JoinMessage) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

// New category input message
type NewCategory struct {
	Name string `json:"name"`
}

// Edit category websocket message
type EditCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Delete category websocket message
type DeleteCategory struct {
	ID string
}

// Category response
type CategoryResponse struct {
	Category models.KanbanCategory `json:"category"`
}

func (m *CategoryResponse) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}
