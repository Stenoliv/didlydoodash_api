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

// IsValidOrganisationRole checks if a string is a valid OrganisationRole
func IsValidOrganisationRole(role string) bool {
	for _, validRole := range OrganisationRoles {
		if strings.EqualFold(string(validRole), role) {
			return true
		}
	}
	return false
}

// ToOrganisationRole converts a string to OrganisationRole, defaults to NotSpecified if invalid
func ToOrganisationRole(role string) OrganisationRole {
	for _, validRole := range OrganisationRoles {
		if strings.EqualFold(string(validRole), role) {
			return validRole
		}
	}
	return NotSpecified
}
