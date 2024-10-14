package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
)

func main() {
	db.Init()

	db.DB.Migrator().DropTable(&models.User{}, &models.UserSession{})
	db.DB.Migrator().DropTable(&models.Project{}, &models.ProjectMember{})
	db.DB.Migrator().DropTable(&models.Organisation{}, &models.OrganisationMember{})
	db.DB.Migrator().DropTable(&models.ChatRoom{}, &models.ChatMember{}, &models.ChatMessage{})
	db.DB.Migrator().DropTable(&models.Kanban{}, &models.KanbanCategory{}, &models.KanbanItem{})
	db.DB.Migrator().DropTable(&models.WhiteboardRoom{}, models.LineData{}, models.LinePoint{})

	// Organisation types
	db.DropType(datatypes.OrganisationRoleName)

	// Project types
	db.DropType(datatypes.ProjectRoleName)
	db.DropType(datatypes.ProjectStatusName)

	// Kanban types
	db.DropType(datatypes.KanbanStatusName)
	db.DropType(datatypes.KanbanItemPriorityName)
}
