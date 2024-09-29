package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
	"fmt"
	"strings"
)

func main() {
	db.Init()

	roles := datatypes.GetOrganisationRolesEnum(datatypes.OrganisationRoles)
	db.CreateType("organisation_role", fmt.Sprintf("ENUM (%s)", roles))

	// Check all tables under user schema
	db.CheckForSchema(strings.Split(datatypes.UserSchema, ".")[0])
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.UserSession{})

	// Check all tables under organisation schema
	db.CheckForSchema(strings.Split(datatypes.OrganisationSchema, ".")[0])
	db.DB.AutoMigrate(&models.Organisation{})
	db.DB.AutoMigrate(&models.OrganisationMember{})
	db.DB.AutoMigrate(&models.ChatRoom{})
	db.DB.AutoMigrate(&models.ChatMember{})
	db.DB.AutoMigrate(&models.ChatMessage{})

	// Check all tables under the project schema
	db.CheckForSchema(strings.Split(datatypes.ProjectSchema, ".")[0])
	db.DB.AutoMigrate(&models.Project{})
	db.DB.AutoMigrate(&models.ProjectMembers{})
}
