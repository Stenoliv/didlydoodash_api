package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func main() {
	db.Init()

	db.DB.Migrator().DropTable(&models.User{}, &models.UserSession{})
	db.DB.Migrator().DropTable(&models.Project{}, &models.ProjectMember{})
	db.DB.Migrator().DropTable(&models.Organisation{}, &models.OrganisationMember{})
	db.DB.Migrator().DropTable(&models.ChatRoom{}, &models.ChatMember{}, &models.ChatMessage{})
	db.DB.Migrator().DropTable(&models.Kanban{}, &models.KanbanCategory{}, &models.KanbanItem{})

	db.DropType("organisation_role")
	db.DropType("project_role")
	db.DropType("project_status")
	db.DropType("kanban_status")
}
