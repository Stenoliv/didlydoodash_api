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

type NewCategoryResponse struct {
	Category models.KanbanCategory `json:"category"`
}

func (m *NewCategoryResponse) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

// Edit category websocket message
type EditCategory struct {
	Name string `json:""`
}

func (m *EditCategory) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}
