package models

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Base struct {
	ID        string     `gorm:"not null;primaryKey;size:21;" json:"id,omitempty"`
	CreatedAt *time.Time `gorm:"deafult:NOW()" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"deafult:NOW()" json:"updatedAt,omitempty"`
}

func (b *Base) GenerateID() (err error) {
	if b.ID == "" {
		id, err := gonanoid.New()
		if err != nil {
			return err
		}
		b.ID = id
	}
	return nil
}

type BaseCreatedIndex struct {
	ID        string     `gorm:"not null;primaryKey;size:21;" json:"id,omitempty"`
	CreatedAt *time.Time `gorm:"deafult:NOW();index;" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"deafult:NOW()" json:"updatedAt,omitempty"`
}

func (b *BaseCreatedIndex) GenerateID() (err error) {
	if b.ID == "" {
		id, err := gonanoid.New()
		if err != nil {
			return err
		}
		b.ID = id
	}
	return nil
}
