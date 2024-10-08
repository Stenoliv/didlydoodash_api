package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetProjects(orgID string) (projects []models.Project, err error) {
	if err := db.DB.Model(&models.Project{}).
		Joins("JOIN project_members ON project_members.project_id = projects.id").
		Where("project_members.user_id = ?", models.CurrentUser).
		Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
