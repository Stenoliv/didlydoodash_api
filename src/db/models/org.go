package models

import (
	"gorm.io/gorm"
)

type Organisation struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	OwnerID   string         `gorm:"not null;size:21;uniqueIndex:idx_o_name_user;" json:"-"`
	Owner     User           `gorm:"not null;" json:"owner"`
	Name      string         `gorm:"not null;uniqueIndex:idx_o_name_user;" json:"name"`
	Chats     []ChatRoom     `gorm:"" json:"chatRooms"`
}

func (o *Organisation) SaveOrganisation(tx *gorm.DB) (err error) {
	if err = tx.Create(&o).Error; err != nil {
		return err
	}
	return nil
}

func (o *Organisation) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.GenerateID()
	if err != nil {
		return err
	}
	return nil
}

func (o *Organisation) AfterFind(tx *gorm.DB) (err error) {
	if o.Owner.ID == "" {
		if err := tx.Model(&User{}).Scopes(PublicUserData).Find(&o.Owner, "id = ?", o.OwnerID).Error; err != nil {
			return err
		}
	}
	return nil
}
