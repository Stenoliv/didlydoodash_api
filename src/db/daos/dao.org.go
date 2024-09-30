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
	if err := db.DB.Model(&models.Organisation{}).
		Where("owner_id = ? OR id IN (SELECT organisation_id FROM organisation_members WHERE user_id = ?)", models.CurrentUser, models.CurrentUser).
		Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}
