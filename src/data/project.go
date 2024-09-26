package data

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"
)

/**
 * Project table
 */
type Project struct {
	Base
	Name           string           `gorm:"size:255;" json:"name"`
	Organisation   Organisation     `gorm:"" json:"organisation"`
	OrganisationID Nanoid           `gorm:"" json:"-"`
	Members        []ProjectMembers `gorm:"" json:"members"`
}

func (p *Project) TableName() string {
	return utils.GetTableName(datatypes.ProjectSchema, p)
}

/**
 * Project members table
 */
type ProjectMembers struct {
	Project   Project `gorm:"" json:"-"`
	ProjectID Nanoid  `gorm:"" json:"-"`
	User      User    `gorm:"" json:"user"`
	UserID    Nanoid  `gorm:"" json:"-"`
}

func (pm *ProjectMembers) TableName() string {
	return utils.GetTableName(datatypes.ProjectSchema, pm)
}
