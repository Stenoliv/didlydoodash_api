package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetAnnouncements(OrganisationID string) ([]models.Announcement, error) {
	var announcement []models.Announcement
	err := db.DB.Model(&models.Announcement{}).Where("organisation_id = ? ", OrganisationID).Find(&announcement).Error

	if err != nil {
		return nil, err
	}
	return announcement, nil
}
func GetAnnouncement(aID string) (*models.Announcement, error) {
	var a *models.Announcement

	err := db.DB.Model(&models.Announcement{}).
		Where("id = ?", aID).First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil

}
