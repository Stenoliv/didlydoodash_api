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
		Joins("JOIN organisation_members ON organisation_members.organisation_id = organisations.id").
		Where("organisation_members.user_id = ?", models.CurrentUser).
		Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}

func GetOrg(id string) (*models.Organisation, error) {
	var org models.Organisation
	if err := db.DB.Model(&models.Organisation{}).Where("id = ?", id).First(&org).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

// Get organistaion members
func GetMembers(id string) (members []models.OrganisationMember, err error) {
	if err = db.DB.Model(&models.OrganisationMember{}).Where("organisation_id = ?", id).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func GetMember(id, userId string) (*models.OrganisationMember, error) {
	var member models.OrganisationMember
	if err := db.DB.Model(&models.OrganisationMember{}).Where("organisation_id = ?", id).Where("user_id = ?", userId).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
