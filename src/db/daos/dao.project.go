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

func GetProject(projectID string) (*models.Project, error) {
	var project models.Project
	if err := db.DB.Model(&models.Project{}).Where("id = ?", projectID).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func GetProjectMember(projectID string, userID string) (*models.ProjectMember, error) {
	var member models.ProjectMember
	if err := db.DB.Model(member).Where("project_id = ? AND user_id = ?", projectID, userID).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
