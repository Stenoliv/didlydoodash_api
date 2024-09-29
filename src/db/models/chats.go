package models

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"

	"gorm.io/gorm"
)

// Chat rooms
type ChatRoom struct {
	Base
	OrganisationID string        `gorm:"size:21;" json:"-"`
	Organisation   Organisation  `gorm:"" json:"-"`
	Name           string        `gorm:"size:255;" json:"name"`
	Members        []ChatMember  `gorm:"" json:"members"`
	Messages       []ChatMessage `gorm:"" json:"messages"`
}

func (o *ChatRoom) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, o)
}

func (o *ChatRoom) BeforeCreate(tx *gorm.DB) error {
	if err := o.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (o *ChatRoom) SaveChatRoom(tx *gorm.DB) error {
	return tx.Create(&o).Error
}

func (o *ChatRoom) AfterFind(tx *gorm.DB) error {
	if err := tx.Association("Members").Find(&o.Members); err != nil {
		return err
	}
	return nil
}

// Member of chats
type ChatMember struct {
	Base
	ChatRoomID string   `gorm:"size:21;uniqueIndex:idx_chat_member;not null;" json:"-"`
	ChatRoom   ChatRoom `gorm:"" json:"-"`
	UserID     string   `gorm:"size:21;uniqueIndex:idx_chat_member;not null;" json:"-"`
	User       User     `gorm:"" json:"user"`
}

func (o *ChatMember) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, o)
}

func (o *ChatMember) AfterFind(tx *gorm.DB) error {
	if err := tx.Association("User").Find(&o.User); err != nil {
		return err
	}
	return nil
}

// Chat message
type ChatMessage struct {
	Base
	ChatRoomID string   `gorm:"size:21;uniqueIndex:idx_room_message;not null;" json:"-"`
	ChatRoom   ChatRoom `gorm:"" json:"-"`
	UserID     string   `gorm:"size:21;uniqueIndex:idx_room_message;not null;" json:"userId"`
	User       User     `gorm:"" json:"-"`
	Message    string   `gorm:"uniqueIndex:idx_room_message;not null;" json:"message"`
}

func (o *ChatMessage) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, o)
}

func (o *ChatMessage) BeforeCreate(tx *gorm.DB) (err error) {
	if err = o.GenerateID(); err != nil {
		return err
	}
	return nil
}
