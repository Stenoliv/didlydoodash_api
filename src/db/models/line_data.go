package models

import (
	"fmt"

	"gorm.io/gorm"
)

type LineData struct {
	Points      []float64 `json:"points"`
	Stroke      string    `json:"stroke"`
	StrokeWidth int       `json:"strokeWidth"`
	Tool        string    `json:"tool"`
	Text        string    `json:"text"`
}

func (l *LineData) SaveLineData(tx *gorm.DB) error { 
	if err := tx.Create(l).Error; err != nil {
		return fmt.Errorf("failed to save line data: %w", err)
	}
	return nil
}