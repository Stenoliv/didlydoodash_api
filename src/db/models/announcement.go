package models

import "gorm.io/gorm"

type Announcement struct {
	Base
	AnnouncmentText string
	OrganisationID  string
}

func (a *Announcement) BeforeCreate(tx *gorm.DB) error {
	err := a.GenerateID()
	if err != nil {
		return err
	}
	return nil
}

func (a *Announcement) SaveAnnouncement(tx *gorm.DB) error {
	return tx.Create(&a).Error
}

func (a *Announcement) BeforeSave(tx *gorm.DB) error {
	return nil
}
