package daos

import (
	"DidlyDoodash-api/src/data"
	"DidlyDoodash-api/src/db"

	"gorm.io/gorm"
)

// Scope functions
func OrganisationListData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "name")
}

// Get all organisations
func GetAllOrgs() (orgs []data.Organisation, err error) {
	if err := db.DB.Model(&data.Organisation{}).Scopes(OrganisationListData).Find(&orgs, "owner_id = ?", data.CurrentUser).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}
