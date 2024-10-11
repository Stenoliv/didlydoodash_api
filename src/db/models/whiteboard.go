package models

import "gorm.io/gorm"

type WhiteboardRoom struct {
	Base
	ProjectID string     `gorm:"size:21;" json:"-"`
	Project   Project    `gorm:"" json:"-"`
	Name      string     `gorm:"size:255;" json:"name"`
	Lines     []LineData `gorm:"-" json:"lines"`
}

func (wb *WhiteboardRoom) BeforeCreate(tx *gorm.DB) error {
	err := wb.GenerateID()
	if err != nil {
		return err
	}
	return nil
}

func (wb *WhiteboardRoom) SaveWhiteboardRoom(tx *gorm.DB) error {
	return tx.Create(&wb).Error
}

func (wb *WhiteboardRoom) BeforeSave(tx *gorm.DB) error {
	return nil
}