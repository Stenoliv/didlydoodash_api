package data

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Base struct {
	ID          string      `gorm:"size:21;primaryKey;" json:"id"`
	CreatedAt *time.Time     `gorm:"deafult:NOW()" json:"createdAt"`
	UpdatedAt *time.Time     `gorm:"deafult:NOW()" json:"updatedAt"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID, err = gonanoid.New()
		if err != nil {
			return err
		}
	}
	now := time.Now()
	if b.CreatedAt == nil {
		b.CreatedAt = &now
	}
	if b.UpdatedAt == nil {
		b.UpdatedAt = &now
	}

	return nil
}