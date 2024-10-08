package models

import (
	"gorm.io/gorm"
)

/**
 * Project table
 */
type Project struct {
	Base
	Name           string           `gorm:"size:255;" json:"name"`
	OrganisationID string           `gorm:"not null;size:21;" json:"-"`
	Organisation   Organisation     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"organisation"`
	Members        []ProjectMembers `gorm:"-" json:"members"`
}

func (o *Project) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.GenerateID()
	if err != nil {
		return err
	}
	return nil
}
