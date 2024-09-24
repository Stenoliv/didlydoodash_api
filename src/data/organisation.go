package data

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"

	"gorm.io/gorm"
)

type Organisation struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	OwnerID   string         `gorm:"not null;" json:"-"`
	Owner     User           `gorm:"not null;" json:"owner"`
	Name      string         `gorm:"not null;" json:"name"`
}

func (o *Organisation) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, o)
}

func (o *Organisation) AfterFind(tx *gorm.DB) (err error) {
	if o.Owner.ID == "" {
		if err := tx.Model(&Organisation{}).Association("Owner").Find(&o.Owner); err != nil {
			return err
		}
	}
	return nil
}

type OrganisationMember struct {
	Base
	OrganisationID string                 `gorm:"uniqueIndex:idx_o_member" json:"-"`
	Organisation   Organisation           `gorm:"" json:"organisation"`
	Role           datatypes.OrganisationRole `gorm:"type:organisation_role" json:"role"`
	UserID         string                 `gorm:"uniqueIndex:idx_o_member" json:"-"`
	User           User                   `gorm:"" json:"user"`
}

func (om *OrganisationMember) TableName() string {
	return utils.GetTableName(datatypes.OrganisationSchema, om)
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
