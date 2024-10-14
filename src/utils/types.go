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
	JoinKanban    MessageType = "kanban.load"
	EditKanban    MessageType = "kanban.edit"
	KanbanError   MessageType = "kanban.error"
	KanbanArchive MessageType = "kanban.archive"
	// Categories
	NewKanbanCategory         MessageType = "kanban.category.new"
	RestoreKanbanCategory     MessageType = "kanban.category.restore"
	EditKanbanCategory        MessageType = "kanban.category.edit"
	DeleteKanbanCategory      MessageType = "kanban.category.delete"
	PermaDeleteKanbanCategory MessageType = "kanban.category.perma"
	// Items
	NewKanbanItem         MessageType = "kanban.item.new"
	RestoreKanbanItem     MessageType = "kanban.item.restore"
	MoveKanbanItem        MessageType = "kanban.item.drop"
	EditKanbanItem        MessageType = "kanban.item.edit"
	DeleteKanbanItem      MessageType = "kanban.item.delete"
	PermaDeleteKanbanItem MessageType = "kanban.item.perma"
)
