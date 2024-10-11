package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
	"fmt"
)

func main() {
	db.Init()

	orgRoles := datatypes.GetOrganisationRolesEnum(datatypes.OrganisationRoles)
	db.CreateType("organisation_role", fmt.Sprintf("ENUM (%s)", orgRoles))
	projectRoles := datatypes.GetProjectRolesEnum(datatypes.ProjectRoles)
	db.CreateType("project_role", fmt.Sprintf("ENUM (%s)", projectRoles))
	projectStatus := datatypes.GetProjectStatusEnum(datatypes.ProjectStatusEnum)
	db.CreateType("project_status", fmt.Sprintf("ENUM (%s)", projectStatus))
	kanbanStatus := datatypes.GetKanbanStatusEnum(datatypes.KanbanStatusEnum)
	db.CreateType("kanban_status", fmt.Sprintf("ENUM (%s)", kanbanStatus))

	// Check all tables for users
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.UserSession{})

	// Check all tables for organisations
	db.DB.AutoMigrate(&models.Organisation{})
	db.DB.AutoMigrate(&models.OrganisationMember{})
	db.DB.AutoMigrate(&models.ChatRoom{})
	db.DB.AutoMigrate(&models.ChatMember{})
	db.DB.AutoMigrate(&models.ChatMessage{})

	// Check all tables for projects
	db.DB.AutoMigrate(&models.Project{})
	db.DB.AutoMigrate(&models.ProjectMember{})

	// Check all tables for kanbans
	db.DB.AutoMigrate(&models.Kanban{})
	db.DB.AutoMigrate(&models.KanbanCategory{})
	db.DB.AutoMigrate(&models.KanbanItem{})
}
