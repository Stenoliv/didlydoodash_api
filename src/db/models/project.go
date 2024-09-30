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
	Organisation   Organisation     `gorm:"" json:"organisation"`
	OrganisationID string           `gorm:"size:21;" json:"-"`
	Members        []ProjectMembers `gorm:"" json:"members"`
}

func (o *Project) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.GenerateID()
	if err != nil {
		return err
	}
	return nil
}
