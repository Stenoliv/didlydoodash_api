package models

import (
	"fmt"

	"gorm.io/gorm"
)

type LineData struct {
	Base
	WhiteboardID string         `gorm:"size:21; not null;" json:"-"`
	Whiteboard   WhiteboardRoom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Points       []LinePoint    `gorm:"not null" json:"points"`
	Stroke       string         `json:"stroke"`
	StrokeWidth  int            `json:"strokeWidth"`
	Tool         string         `json:"tool"`
	Text         string         `json:"text"`
}

func (l *LineData) SaveLineData(tx *gorm.DB) error {
	if err := tx.Create(l).Error; err != nil {
		return fmt.Errorf("failed to save line data: %w", err)
	}
	return nil
}
func (ld *LineData) BeforeCreate(tx *gorm.DB) error {
	err := ld.GenerateID()
	if err != nil {
		return err
	}
	return nil
}
