package data

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Base struct {
	ID        string     `gorm:"not null;size:21;primaryKey;unqiue;" json:"id"`
	CreatedAt *time.Time `gorm:"deafult:NOW()" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"deafult:NOW()" json:"updatedAt"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID, err = gonanoid.New()
		if err != nil {
			return err
		}
	}
	return nil
}
