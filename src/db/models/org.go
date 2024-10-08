package models

import (
	"DidlyDoodash-api/src/db/datatypes"

	"gorm.io/gorm"
)

type Organisation struct {
	Base
	Owner   OrganisationMember   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"owner"`
	Name    string               `gorm:"not null;" json:"name"`
	Chats   []ChatRoom           `gorm:"-" json:"chatRooms,omitempty"`
	Members []OrganisationMember `gorm:"-" json:"members,omitempty"`
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
	if o.Owner.User.ID == "" {
		if err := tx.Model(&OrganisationMember{}).Where("organisation_id = ?", o.ID).Where("role = ?", datatypes.CEO).First(&o.Owner).Error; err != nil {
			return err
		}
	}
	return nil
}
