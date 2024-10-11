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

// Kanban Status
type KanbanStatus string

const (
	KANBAN_PLANNING    KanbanStatus = "Planning"
	KANBAN_IN_PROGRESS KanbanStatus = "In Progress"
	KANBAN_DONE        KanbanStatus = "Done"
	KANBAN_ARCHIVED    KanbanStatus = "Archived"
)

var KanbanStatusEnum = []KanbanStatus{
	KANBAN_PLANNING,
	KANBAN_IN_PROGRESS,
	KANBAN_DONE,
	KANBAN_ARCHIVED,
}

// GetKanbanStatusEnum returns a string representation of the valid Kanban statuses
func GetKanbanStatusEnum(statuses []KanbanStatus) string {
	var statusStrings []string
	for _, status := range statuses {
		statusStrings = append(statusStrings, fmt.Sprintf("'%s'", status))
	}
	return strings.Join(statusStrings, ", ")
}

// IsValidKanbanStatus checks if the provided status is a valid Kanban status
func IsValidKanbanStatus(status string) bool {
	for _, validStatus := range KanbanStatusEnum {
		if strings.EqualFold(string(validStatus), status) {
			return true
		}
	}
	return false
}

// ToKanbanStatus converts a string to a KanbanStatus type, returning PLANNING if invalid
func ToKanbanStatus(status string) KanbanStatus {
	for _, validStatus := range KanbanStatusEnum {
		if strings.EqualFold(string(validStatus), status) {
			return validStatus
		}
	}
	return KANBAN_PLANNING
}
