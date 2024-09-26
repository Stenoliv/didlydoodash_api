package data

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Base struct {
	ID        Nanoid     `gorm:"primaryKey;" json:"id"`
	CreatedAt *time.Time `gorm:"deafult:NOW()" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"deafult:NOW()" json:"updatedAt"`
}

type Nanoid string

func (n *Nanoid) GormDataType() string {
	return "VARCHAR(21)"
}

func (n *Nanoid) BeforeCreate(tx *gorm.DB) (err error) {
	if *n == "" {
		id, err := gonanoid.New()
		if err != nil {
			return err
		}
		*n = Nanoid(id)
	}
	return nil
}
