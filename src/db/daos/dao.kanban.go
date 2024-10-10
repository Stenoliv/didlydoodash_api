package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetAllKanbans(id string) (kanbans []models.Kanban, err error) {
	if err := db.DB.Model(&models.Kanban{}).Where("project_id = ?", id).Scan(&kanbans).Error; err != nil {
		return nil, err
	}
	return kanbans, nil
}
