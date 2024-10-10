package models

import (
	"DidlyDoodash-api/src/db/datatypes"

	"gorm.io/gorm"
)

/**
 * Project table
 */
type Project struct {
	Base
	Name           string                  `gorm:"size:255;" json:"name"`
	OrganisationID string                  `gorm:"not null;size:21;" json:"-"`
	Organisation   Organisation            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"organisation"`
	Status         datatypes.ProjectStatus `gorm:"type:project_status;" json:"status"`
	Members        []ProjectMember         `gorm:"-" json:"members"`
}

func (o *Project) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.GenerateID()
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) AfterFind(tx *gorm.DB) (err error) {
	if err := tx.Model(&ProjectMember{}).Where("project_id = ?", p.ID).Find(&p.Members).Error; err != nil {
		return err
	}
	return nil
}

func (p *Project) SaveProject(tx *gorm.DB) (err error) {
	return tx.Create(&p).Error
}
