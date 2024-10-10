package datatypes

import (
	"fmt"
	"strings"
)

type OrganisationRole string

const (
	// Executive & Leadership Roles
	CEO OrganisationRole = "CEO"

	// Management Roles
	ProjectManager OrganisationRole = "Project Manager"
	HRManager      OrganisationRole = "HR Manager"
	ITManager      OrganisationRole = "IT Manager"

	// Technical Roles
	SeniorSoftwareEngineer OrganisationRole = "Senior Software Engineer"
	JuniorSoftwareEngineer OrganisationRole = "Junior Software Engineer"
	ITSupport              OrganisationRole = "IT Support"

	// Human Resources & Administration Roles
	Recruiter OrganisationRole = "Recruiter"

	// Role not specified
	NotSpecified OrganisationRole = "Not specified"
)

var OrganisationRoles = []OrganisationRole{
	CEO,
	ProjectManager,
	ITManager,
	SeniorSoftwareEngineer,
	JuniorSoftwareEngineer,
	ITSupport,
	HRManager,
	Recruiter,
	NotSpecified,
}

func GetOrganisationRolesEnum(roles []OrganisationRole) string {
	var roleStrings []string
	for _, role := range roles {
		roleStrings = append(roleStrings, fmt.Sprintf("'%s'", role))
	}
	return strings.Join(roleStrings, ", ")
}

func IsValidOrganisationRole(role string) bool {
	for _, validRole := range OrganisationRoles {
		if strings.EqualFold(string(validRole), role) {
			return true
		}
	}
	return false
}

func ToOrganisationRole(role string) OrganisationRole {
	for _, validRole := range OrganisationRoles {
		if strings.EqualFold(string(validRole), role) {
			return validRole
		}
	}
	return NotSpecified
}

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
