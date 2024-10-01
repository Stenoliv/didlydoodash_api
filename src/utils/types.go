package utils

type Tokens struct {
	Access  *string `json:"access"`
	Refresh *string `json:"refresh"`
}

type MessageType string

const (
	MessageSend   MessageType = "message.send"
	SendMessages  MessageType = "message.all"
	MessageTyping MessageType = "message.type"
	MessageError  MessageType = "message.error"
)
