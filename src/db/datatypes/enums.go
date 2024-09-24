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
	GeneralManager   OrganisationRole = "General Manager"
	ProjectManager   OrganisationRole = "Project Manager"
	SalesManager     OrganisationRole = "Sales Manager"
	MarketingManager OrganisationRole = "Marketing Manager"
	HRManager        OrganisationRole = "HR Manager"
	ITManager        OrganisationRole = "IT Manager"

	// Technical Roles
	SoftwareEngineer    OrganisationRole = "Software Engineer"
	ITSupportSpecialist OrganisationRole = "IT Support Specialist"

	// Marketing & Sales Roles
	AccountManager     OrganisationRole = "Account Manager"
	SocialMediaManager OrganisationRole = "Social Media Manager"

	// Finance & Accounting Roles
	Accountant OrganisationRole = "Accountant"

	// Human Resources & Administration Roles
	Recruiter    OrganisationRole = "Recruiter"
	HRGeneralist OrganisationRole = "HR Generalist"

	// Operations & Logistics Roles
	WarehouseManager     OrganisationRole = "Warehouse Manager"
	LogisticsCoordinator OrganisationRole = "Logistics Coordinator"

	// Customer Support Roles
	HelpDeskTechnician OrganisationRole = "Help Desk Technician"
)

var OrganisationRoles = []OrganisationRole{
	CEO,
	GeneralManager,
	ProjectManager,
	SalesManager,
	MarketingManager,
	HRManager,
	ITManager,
	SoftwareEngineer,
	ITSupportSpecialist,
	AccountManager,
	SocialMediaManager,
	Accountant,
	Recruiter,
	HRGeneralist,
	WarehouseManager,
	LogisticsCoordinator,
	HelpDeskTechnician,
}

func GetOrganisationRolesEnum(roles []OrganisationRole) string {
	var roleStrings []string
	for _, role := range roles {
		roleStrings = append(roleStrings, fmt.Sprintf("'%s'", role))
	}
	return strings.Join(roleStrings, ", ")
}