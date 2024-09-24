package data

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"

	"gorm.io/gorm"
)

type User struct {
	Base
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Username   string         `gorm:"size:50;not null:unique;index;" json:"username"`
	Email      string         `gorm:"size:255;not null;unique;" json:"email"`
	Password   string         `gorm:"size:255;not null;" json:"-"`
	IsVerified bool           `gorm:"default:false;" json:"isVerified,omitempty"`
}

func (u *User) TableName() string {
	return utils.GetTableName(datatypes.UserSchema, u)
}
