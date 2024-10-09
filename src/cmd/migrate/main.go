package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
	"fmt"
)

func main() {
	db.Init()

	roles := datatypes.GetOrganisationRolesEnum(datatypes.OrganisationRoles)
	db.CreateType("organisation_role", fmt.Sprintf("ENUM (%s)", roles))

	// Check all tables for users
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.UserSession{})

	// Check all tables for organisations
	db.DB.AutoMigrate(&models.Organisation{})
	db.DB.AutoMigrate(&models.OrganisationMember{})
	db.DB.AutoMigrate(&models.ChatRoom{})
	db.DB.AutoMigrate(&models.ChatMember{})
	db.DB.AutoMigrate(&models.ChatMessage{})
	db.DB.AutoMigrate(&models.ChatMessageReadStatus{})

	// Check all tables for projects
	db.DB.AutoMigrate(&models.Project{})
	db.DB.AutoMigrate(&models.ProjectMembers{})
}
