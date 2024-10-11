package utils

type Tokens struct {
	Access  *string `json:"access"`
	Refresh *string `json:"refresh"`
}

type MessageType string

const (
	MessageSend      MessageType = "message.send"
	MessageRead      MessageType = "message.read"
	SendMessages     MessageType = "message.all"
	LoadMoreMessages MessageType = "message.more"
	MessageTyping    MessageType = "message.type"
	MessageError     MessageType = "message.error"

	// Kanban messages
	JoinKanban           MessageType = "kanban.load"
	NewKanbanCategory    MessageType = "kanban.category.new"
	EditKanbanCategory   MessageType = "kanban.category.edit"
	DeleteKanbanCategory MessageType = "kanban.category.delete"
)
