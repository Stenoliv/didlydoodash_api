package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"

	"gorm.io/gorm"
)

// Scope functions
func OrganisationListData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "name")
}

// Get all organisations
func GetAllOrgs() (orgs []models.Organisation, err error) {
	if err := db.DB.Model(&models.Organisation{}).Scopes(OrganisationListData).Find(&orgs, "owner_id = ?", models.CurrentUser).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}
