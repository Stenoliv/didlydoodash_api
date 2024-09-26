package main

import (
	"DidlyDoodash-api/src/data"
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"fmt"
	"strings"
)

func main() {
	db.Init()

	roles := datatypes.GetOrganisationRolesEnum(datatypes.OrganisationRoles)
	db.CreateType("organisation_role", fmt.Sprintf("ENUM (%s)", roles))

	// Check all tables under user schema
	db.CheckForSchema(strings.Split(datatypes.UserSchema, ".")[0])
	db.DB.AutoMigrate(&data.User{})

	// Check all tables under organisation schema
	db.CheckForSchema(strings.Split(datatypes.OrganisationSchema, ".")[0])
	db.DB.AutoMigrate(&data.Organisation{})
	db.DB.AutoMigrate(&data.OrganisationMember{})

	// Check all tables under the project schema
	db.CheckForSchema(strings.Split(datatypes.ProjectSchema, ".")[0])
	db.DB.AutoMigrate(&data.Project{})
	db.DB.AutoMigrate(&data.ProjectMembers{})
}
