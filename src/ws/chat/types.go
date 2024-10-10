package chat

import "encoding/json"

type MessageSend struct {
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type MessageRead struct {
	MessageID string `json:"messageId" binding:"required"`
}

type NewMessage struct {
	ChatID  string `json:"chatId" binding:"required"`
	UserID  string `json:"userId" binding:"required"`
	Message string `json:"message"`
}

func (cm *NewMessage) ToJSON() ([]byte, error) {
	return json.Marshal(&cm)
}

type CountMessage struct {
	ChatID string `json:"chatId"`
	UserID string `json:"userId"`
	Count  int64  `json:"unreadMessages"`
}

func (cm *CountMessage) ToJSON() ([]byte, error) {
	return json.Marshal(&cm)
}

var (
	MessageCount string = "message.count"
	MessageNew   string = "message.new"
)
