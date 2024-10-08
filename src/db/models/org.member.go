package models

import (
	"DidlyDoodash-api/src/db/datatypes"

	"gorm.io/gorm"
)

// Organisation member
type OrganisationMember struct {
	OrganisationID string                     `gorm:"not null;uniqueIndex:idx_o_member;size:21;" json:"-"`
	Organisation   *Organisation              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Role           datatypes.OrganisationRole `gorm:"type:organisation_role" json:"role"`
	UserID         string                     `gorm:"not null;uniqueIndex:idx_o_member;size:21;" json:"-"`
	User           User                       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (om *OrganisationMember) SaveMember(db *gorm.DB) (err error) {
	return db.Create(&om).Error
}

func (om *OrganisationMember) AfterCreate(tx *gorm.DB) (err error) {
	if om.User.ID == "" {
		if err := tx.Model(&User{}).Scopes(PublicUserData).Where("id = ?", om.UserID).Find(&om.User).Error; err != nil {
			return err
		}
	}
	return nil
}

func (om *OrganisationMember) AfterFind(tx *gorm.DB) (err error) {
	if om.User.ID == "" {
		if err := tx.Model(&User{}).Scopes(PublicUserData).Where("id = ?", om.UserID).Find(&om.User).Error; err != nil {
			return err
		}
	}
	return nil
}
