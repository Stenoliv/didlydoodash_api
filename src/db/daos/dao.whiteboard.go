package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetWhiteboards(prID string) ([]models.WhiteboardRoom, error) {
	var whiteboardRooms []models.WhiteboardRoom
	err := db.DB.Model(&models.WhiteboardRoom{}).Where("project_id = ? AND EXISTS (SELECT 1 FROM project_members WHERE user_id = ? AND project_id = ?)", prID, models.CurrentUser, prID).Find(&whiteboardRooms).Error

	if err != nil {
		return nil, err
	}
	return whiteboardRooms, nil
}

func GetWhiteboard(wbID string) (*models.WhiteboardRoom, error) {
	var wb *models.WhiteboardRoom

	err := db.DB.Model(&models.WhiteboardRoom{}).
		Where("id = ?", wbID).First(&wb).Error
	if err != nil {
		return nil, err
	}
	return wb, nil

}
