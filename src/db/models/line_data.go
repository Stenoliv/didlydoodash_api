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
func (l *LinePoint) SaveLinePoint(tx *gorm.DB) error {
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
func (ld *LineData) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&LinePoint{}).Where("line_data_id = ?", ld.ID).Order("created_at ASC").Find(&ld).Error; err != nil {
		return err
	}
	return nil
}
