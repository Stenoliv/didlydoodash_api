package models

import (
	"fmt"

	"gorm.io/gorm"
)

type LinePoint struct {
	Point      float64  `json:"point"`
	LineDataID string   `gorm:"size:21; not null;" json:"-"`
	LineData   LineData `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (l *LinePoint) SaveLinePoint(tx *gorm.DB) error {
	if err := tx.Create(l).Error; err != nil {
		return fmt.Errorf("failed to save line data: %w", err)
	}
	return nil
}
func (ld *LineData) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&LinePoint{}).Where("line_data_id = ?", ld.ID).Find(&ld.Points).Error; err != nil {
		return err
	}
	return nil
}
