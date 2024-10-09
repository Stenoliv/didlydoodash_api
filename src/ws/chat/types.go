package chat

type MessageSend struct {
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type MessageRead struct {
	MessageID string `json:"messageId" binding:"required"`
}
