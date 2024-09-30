package models

import (
	"DidlyDoodash-api/src/db/datatypes"

	"gorm.io/gorm"
)

// Organisation member
type OrganisationMember struct {
	OrganisationID string                     `gorm:"not null;uniqueIndex:idx_o_member;size:21;" json:"-"`
	Organisation   Organisation               `gorm:"" json:"organisation"`
	Role           datatypes.OrganisationRole `gorm:"type:organisation_role" json:"role"`
	UserID         string                     `gorm:"not null;uniqueIndex:idx_o_member;size:21;" json:"-"`
	User           User                       `gorm:"" json:"user"`
}

func (om *OrganisationMember) SaveMember(db *gorm.DB) (err error) {
	return db.Create(&om).Error
}

func (om *OrganisationMember) AfterFind(tx *gorm.DB) (err error) {
	if om.Organisation.ID == "" {
		if err := tx.Model(&OrganisationMember{}).Association("Organisation").Find(&om.Organisation); err != nil {
			return err
		}
	}
	if om.User.ID == "" {
		if err := tx.Model(&OrganisationMember{}).Association("User").Find(&om.Organisation); err != nil {
			return err
		}
	}
	return nil
}
