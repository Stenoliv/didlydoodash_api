package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func main() {
	db.Init()

	db.DB.Migrator().DropTable(&models.User{}, &models.UserSession{})
	db.DB.Migrator().DropTable(&models.Project{}, &models.ProjectMembers{})
	db.DB.Migrator().DropTable(&models.Organisation{}, &models.OrganisationMember{})
	db.DB.Migrator().DropTable(&models.ChatRoom{}, &models.ChatMember{}, &models.ChatMessage{})
}
