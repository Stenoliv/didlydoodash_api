package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetWhiteboards(prID string) ([]models.WhiteboardRoom , error) {
	var whiteboardRooms []models.WhiteboardRoom
	err := db.DB.Model(&models.WhiteboardRoom{}).Where("project_id = ?", prID).Find(&whiteboardRooms).Error 

	if err != nil {
		return nil, err
	}
	return whiteboardRooms, nil
}

func GetWhiteboard(wbID string) (*models.WhiteboardRoom, error){
	var wb models.WhiteboardRoom

	err := db.DB.Model(&models.WhiteboardRoom{}).
		Find(&wb, "id = ?", wbID).Error
	if err != nil {
		return nil, err
	}
	return &wb, nil

}

