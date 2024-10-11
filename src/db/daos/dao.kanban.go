package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetAllKanbans(id string) (kanbans []models.Kanban, err error) {
	if err := db.DB.Model(&models.Kanban{}).Where("project_id = ?", id).Order("created_at ASC").Find(&kanbans).Error; err != nil {
		return nil, err
	}
	return kanbans, nil
}

func GetKanban(id string) (*models.Kanban, error) {
	var kanban *models.Kanban
	if err := db.DB.Model(&models.Kanban{}).Where("id = ?", id).Find(&kanban).Error; err != nil {
		return nil, err
	}
	return kanban, nil
}
