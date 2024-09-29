package models

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"
)

/**
 * Project members table
 */
type ProjectMembers struct {
	Project   Project `gorm:"" json:"-"`
	ProjectID string  `gorm:"size:21;" json:"-"`
	User      User    `gorm:"" json:"user"`
	UserID    string  `gorm:"size:21;" json:"-"`
}

func (pm *ProjectMembers) TableName() string {
	return utils.GetTableName(datatypes.ProjectSchema, pm)
}
