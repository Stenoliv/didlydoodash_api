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
	db.DB.AutoMigrate(&models.User{}, &models.UserSession{})

	// Check all tables for organisations
	db.DB.AutoMigrate(&models.Organisation{}, &models.OrganisationMember{})
	db.DB.AutoMigrate(&models.ChatRoom{}, &models.ChatMember{}, &models.ChatMessage{})

	// Check all tables for projects
	db.DB.AutoMigrate(&models.Project{}, &models.ProjectMember{})

	// Check all tables for kanbans
	db.DB.AutoMigrate(&models.Kanban{}, &models.KanbanCategory{}, &models.KanbanItem{})

	db.DB.AutoMigrate(&models.WhiteboardRoom{})
	db.DB.AutoMigrate(&models.LineData{})
	db.DB.AutoMigrate(&models.LinePoint{})

	db.DB.AutoMigrate(&models.Announcement{})
}
