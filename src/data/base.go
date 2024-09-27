package data

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Base struct {
	ID        string     `gorm:"primaryKey;" json:"id"`
	CreatedAt *time.Time `gorm:"deafult:NOW()" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"deafult:NOW()" json:"updatedAt"`
}
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
		id, err := gonanoid.New()
		if err != nil {
			return err
		}
		b.ID = id
	return nil
}
type Nanoid string

func (n *Nanoid) GormDataType() string {
	return "VARCHAR(21)"
}


