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

// Get organistaion members
func GetMembers(id string) (members []models.OrganisationMember, err error) {
	if err = db.DB.Model(&models.OrganisationMember{}).Where("organisation_id = ?", id).Where("user_id != ?", models.CurrentUser).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func GetMember(id, userId string) (*models.OrganisationMember, error) {
	var member models.OrganisationMember
	if err := db.DB.Model(&models.OrganisationMember{}).Where("organisation_id = ?", id).Where("user_id = ?", userId).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
