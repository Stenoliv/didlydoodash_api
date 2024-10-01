package chat

type MessageStruct struct {
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}
