package models

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"

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

func (p *Project) TableName() string {
	return utils.GetTableName(datatypes.ProjectSchema, p)
}

func (o *Project) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.GenerateID()
	if err != nil {
		return err
	}
	return nil
}
