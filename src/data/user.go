package data

import "gorm.io/gorm"

type User struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Username string `gorm:"size:25;not null:unique;" json:"username"`
	Email string `gorm:"size:255;not null;unique;" json:"email"`
	Password string `gorm:"size:255;not null;unique" json:"-"`
	IsVerified bool `gorm:"default:false;" json:"isVerified,omitempty"`
}