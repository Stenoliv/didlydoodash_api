package data

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"

	"gorm.io/gorm"
)

type Organisation struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	OwnerID   string         `gorm:"not null;size:21;uniqueIndex:idx_o_name_user;" json:"-"`
	Owner     User           `gorm:"not null;" json:"owner"`
	Name      string         `gorm:"not null;uniqueIndex:idx_o_name_user;" json:"name"`
}

func (o *Organisation) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, o)
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
	if o.Owner.ID == "" {
		if err := tx.Model(&Organisation{}).Association("Owner").Find(&o.Owner); err != nil {
			return err
		}
	}
	return nil
}

// Organisation member
type OrganisationMember struct {
	OrganisationID string                     `gorm:"not null;uniqueIndex:idx_o_member;size:21;" json:"-"`
	Organisation   Organisation               `gorm:"" json:"organisation"`
	Role           datatypes.OrganisationRole `gorm:"type:organisation_role" json:"role"`
	UserID         string                     `gorm:"not null;uniqueIndex:idx_o_member;size:21;" json:"-"`
	User           User                       `gorm:"" json:"user"`
}

func (om *OrganisationMember) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, om)
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
