package models

import (
	"DidlyDoodash-api/src/db/datatypes"

	"gorm.io/gorm"
)

/**
 * Project members table
 */
type ProjectMember struct {
	ProjectID string                `gorm:"size:21;uniqueIndex:idx_projectMember;" json:"projectId"`
	Project   Project               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID    string                `gorm:"size:21;uniqueIndex:idx_projectMember;" json:"-"`
	User      User                  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Role      datatypes.ProjectRole `gorm:"type:project_role;not null;" json:"role"`
}

func (pm *ProjectMember) SaveMember(tx *gorm.DB) error {
	return tx.Create(&pm).Error
}

func (pm *ProjectMember) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", pm.UserID).Find(&pm.User).Error; err != nil {
		return err
	}
	return nil
}
