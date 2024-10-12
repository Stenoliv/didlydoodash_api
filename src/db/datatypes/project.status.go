package datatypes

import (
	"fmt"
	"strings"
)

const (
	ProjectStatusName string = "project_status"
)

// Project Status
type ProjectStatus string

const (
	ACTIVE    ProjectStatus = "Active"
	COMPLETED ProjectStatus = "Completed"
	ARCHIVED  ProjectStatus = "Archived"
)

var ProjectStatusEnum = []ProjectStatus{
	ACTIVE,
	COMPLETED,
	ARCHIVED,
}

func GetProjectStatusEnum(roles []ProjectStatus) string {
	var roleStrings []string
	for _, role := range roles {
		roleStrings = append(roleStrings, fmt.Sprintf("'%s'", role))
	}
	return strings.Join(roleStrings, ", ")
}

func IsValidProjectStatus(role string) bool {
	for _, validRole := range ProjectStatusEnum {
		if strings.EqualFold(string(validRole), role) {
			return true
		}
	}
	return false
}

func ToProjectStatus(role string) ProjectStatus {
	for _, validRole := range ProjectStatusEnum {
		if strings.EqualFold(string(validRole), role) {
			return validRole
		}
	}
	return ACTIVE
}
