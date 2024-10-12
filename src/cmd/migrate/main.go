package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
	"fmt"
)

func main() {
	db.Init()

	// Organisation types
	orgRoles := datatypes.GetOrganisationRolesEnum(datatypes.OrganisationRoles)
	db.CreateType(datatypes.OrganisationRoleName, fmt.Sprintf("ENUM (%s)", orgRoles))

	// Project types
	projectRoles := datatypes.GetProjectRolesEnum(datatypes.ProjectRoles)
	db.CreateType(datatypes.ProjectRoleName, fmt.Sprintf("ENUM (%s)", projectRoles))
	projectStatus := datatypes.GetProjectStatusEnum(datatypes.ProjectStatusEnum)
	db.CreateType(datatypes.ProjectStatusName, fmt.Sprintf("ENUM (%s)", projectStatus))

	// Kanban types
	kanbanStatus := datatypes.GetKanbanStatusEnum(datatypes.KanbanStatusEnum)
	db.CreateType(datatypes.KanbanStatusName, fmt.Sprintf("ENUM (%s)", kanbanStatus))
	KanbanItemPriority := datatypes.GetKanbanItemPriorityEnum(datatypes.KanbanItemPriorityEnum)
	db.CreateType(datatypes.KanbanItemPriorityName, fmt.Sprintf("ENUM (%s)", KanbanItemPriority))

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
