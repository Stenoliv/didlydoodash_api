package models

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"
	"time"

	"gorm.io/gorm"
)

// Store refresh tokens of user in a seperate table
type UserSession struct {
	UserID     string     `gorm:"size:21;uniquIndex:idx_token;" json:"-"`
	User       User       `gorm:"" json:"-"`
	JTI        string     `gorm:"size:21;unique;uniquIndex:idx_token;" json:"-"`
	ExpireDate *time.Time `gorm:"not null;uniquIndex:idx_token;" json:"-"`
	RememberMe bool       `gorm:"not null;" json:"-" default:"false"`
}

func (us *UserSession) TableName() string {
	return utils.GetTableName(datatypes.UserSchema, us)
}

func (o *UserSession) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (o *UserSession) SaveSession(tx *gorm.DB) (err error) {
	if err = tx.Create(o).Error; err != nil {
		return err
	}
	return nil
}
