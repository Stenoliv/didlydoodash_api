package data

import (
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var CurrentUser *User

// User table struct
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

func (u *User) SaveUser(tx *gorm.DB) error {
	if err := tx.Create(&u).Error; err != nil {
		return err
	}
	return nil
}
func (u *User) BeforeCreate(tx *gorm.DB) error {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	return nil
} 
func (u *User)Validatepassword(input string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(u.Password) , []byte(input) )
	return err == nil	

}
// Store refresh tokens of user in a seperate table
type UserSessions struct {
	Base
	UserID     Nanoid     `gorm:"" json:"-"`
	User       User       `gorm:"" json:"-"`
	JTI        Nanoid     `gorm:"" json:"-"`
	ExpireDate *time.Time `gorm:"" json:"-"`
	RememberMe bool       `gorm:"default:false;" json:"-"`
}

func (us *UserSessions) TableName() string {
	return utils.GetTableName(datatypes.UserSchema, us)
}
