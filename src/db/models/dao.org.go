package models

import (
	"DidlyDoodash-api/src/db"

	"gorm.io/gorm"
)

// Scope functions
func OrganisationListData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "name")
}

// Get all organisations
func GetAllOrgs() (orgs []Organisation, err error) {
	if err := db.DB.Model(&Organisation{}).Scopes(OrganisationListData).Find(&orgs, "owner_id = ?", CurrentUser).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}
