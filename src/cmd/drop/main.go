package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"strings"
)

func main() {
	db.Init()

	db.DropSchema(strings.Split(datatypes.UserSchema, ".")[0])
	db.DropSchema(strings.Split(datatypes.OrganisationSchema, ".")[0])
	// db.DropSchema(strings.Split(datatypes.ProjectSchema, ".")[0])
}
