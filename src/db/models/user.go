package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var CurrentUser *string

func PublicUserData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "username")
}

// User table struct
type User struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Username  string         `gorm:"size:50;not null:unique;index;" json:"username,omitempty"`
	Email     string         `gorm:"size:255;not null;unique;" json:"email,omitempty"`
	Password  string         `gorm:"size:255;not null;" json:"-"`
}

func (u *User) SaveUser(tx *gorm.DB) error {
	if err := tx.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (o *User) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.GenerateID()
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(o.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	o.Password = string(hash)
	return nil
}

func (u *User) Validatepassword(input string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input)) == nil
}
