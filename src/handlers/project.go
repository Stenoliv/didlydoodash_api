package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOrganisation(c *gin.Context) *models.Organisation {
	org, exists := c.Get("organisation")
	if !exists {
		return nil
	}

	organisation, ok := org.(*models.Organisation)
	if !ok {
		return nil
	}

	return organisation
}

func GetAllProjects(c *gin.Context) {
	// Get organisation from context
	org := getOrganisation(c)
	if org == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.OrgNotFound)
		return
	}

	// Get all projects
	projects, err := daos.GetProjects(org.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.OrgNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// Create new project
type createProjectInput struct {
	Name    string               `json:"name" binding:"required"`
	Members []projectMemberInput `json:"members"`
}

type projectMemberInput struct {
	User models.User `json:"user"`
	Role string      `json:"role"`
}

func CreateProjects(c *gin.Context) {
	var input createProjectInput
	// Bind request body to input struct
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	// Get organisation from context
	org := getOrganisation(c)
	if org == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.OrgNotFound)
		return
	}

	// Start transaction
	tx := db.DB.Begin()

	project := &models.Project{
		Name:           input.Name,
		Status:         datatypes.ACTIVE,
		OrganisationID: org.ID,
	}

	// Save project to transaction
	if err := project.SaveProject(tx); err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ProjectCreateError)
		return
	}

	for _, member := range input.Members {
		projectMember := &models.ProjectMember{
			UserID:    member.User.ID,
			ProjectID: project.ID,
			Project:   *project,
			Role:      datatypes.ToProjectRole(member.Role),
		}

		if err := projectMember.SaveMember(tx); err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ProjectCreateError)
			return
		}

		project.Members = append(project.Members, *projectMember)
	}

	// Try to commit to database
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

func UpdateProjects(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

func DeleteProjects(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

func PermaDeleteProjects(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

/**
 * Project members
 */
func GetProjectMembers(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
func GetProjectMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
func UpdateProjectMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
func DeleteProjectMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
