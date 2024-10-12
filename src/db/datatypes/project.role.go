package datatypes

import (
	"fmt"
	"strings"
)

const (
	ProjectRoleName string = "project_role"
)

// Project Role
type ProjectRole string

const (
	ADMIN ProjectRole = "Admin"
	EDIT  ProjectRole = "Edit"
	VIEW  ProjectRole = "View"
)

var ProjectRoles = []ProjectRole{
	ADMIN,
	EDIT,
	VIEW,
}

func GetProjectRolesEnum(roles []ProjectRole) string {
	var roleStrings []string
	for _, role := range roles {
		roleStrings = append(roleStrings, fmt.Sprintf("'%s'", role))
	}
	return strings.Join(roleStrings, ", ")
}

func IsValidProjectRole(role string) bool {
	for _, validRole := range ProjectRoles {
		if strings.EqualFold(string(validRole), role) {
			return true
		}
	}
	return false
}

func ToProjectRole(role string) ProjectRole {
	for _, validRole := range ProjectRoles {
		if strings.EqualFold(string(validRole), role) {
			return validRole
		}
	}
	return VIEW
}
