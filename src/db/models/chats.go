package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Chat rooms
type ChatRoom struct {
	Base
	OrganisationID string        `gorm:"size:21;" json:"-"`
	Organisation   Organisation  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Name           string        `gorm:"size:255;" json:"name"`
	Members        []ChatMember  `gorm:"-" json:"members"`
	Messages       []ChatMessage `gorm:"-" json:"messages"`
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
	if o.Members == nil {
		if err := tx.Model(&ChatMember{}).Where("room_id = ?", o.ID).Find(&o.Members).Error; err != nil {
			return err
		}
	}
	return nil
}

// Member of chats
type ChatMember struct {
	Base
	RoomID string   `gorm:"size:21;not null;" json:"-"`
	Room   ChatRoom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID string   `gorm:"size:21;not null;" json:"-"`
	User   User     `gorm:"" json:"member"`
}

func (o *ChatMember) SaveMember(tx *gorm.DB) error {
	return tx.Create(&o).Error
}

func (o *ChatMember) BeforeCreate(tx *gorm.DB) error {
	if err := o.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (o *ChatMember) AfterCreate(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", o.UserID).Find(&o.User).Error; err != nil {
		return err
	}
	return nil
}

func (o *ChatMember) AfterFind(tx *gorm.DB) (err error) {
	if o.User.ID == "" {
		if err = tx.Model(&User{}).Scopes(PublicUserData).Where("id = ?", o.UserID).Find(&o.User).Error; err != nil {
			return err
		}
	}
	return nil
}

// Chat message
type ChatMessage struct {
	Base
	RoomID  string   `gorm:"size:21;not null;" json:"-"`
	Room    ChatRoom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID  string   `gorm:"size:21;not null;" json:"userId"`
	User    User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Message string   `gorm:"not null;" json:"message"`
}

func (o *ChatMessage) BeforeCreate(tx *gorm.DB) (err error) {
	if err = o.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (o *ChatMessage) SaveMessage(tx *gorm.DB) error {
	return tx.Create(&o).Error
}

func (o *ChatMessage) ToJSON() []byte {
	data, err := json.Marshal(o)
	if err != nil {
		return nil
	}
	return data
}

type ChatMessageReadStatus struct {
	Base
	MessageID string      `gorm:"size:21;not null;" json:"-"` // The message ID
	Message   ChatMessage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID    string      `gorm:"size:21;not null;" json:"userId"` // The user who read the message
	User      User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ReadAt    time.Time   `gorm:"not null;" json:"readAt"` // The timestamp when the message was read
}
