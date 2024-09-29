package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/datatypes"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrganisations(c *gin.Context) {
	orgs, err := models.GetAllOrgs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"organisations": orgs})
}

type OrganisationInput struct {
	Name    string                    `json:"name" binding:"required"`
	Members []OrganisationMemberInput `json:"members"`
}

type OrganisationMemberInput struct {
	ID   string                     `json:"userId"`
	Role datatypes.OrganisationRole `json:"role"`
}

// Create organisation function
func CreateOrganisation(c *gin.Context) {
	var input OrganisationInput
	tx := db.DB.Begin()
	// Bind request body to input struct
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Create Organisation object
	org := &models.Organisation{
		Name:    input.Name,
		OwnerID: *models.CurrentUser,
	}

	// Save organisation object to database
	if err := org.SaveOrganisation(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.FailedToCreateOrg)
		return
	}

	// Create organisation members
	for _, member := range input.Members {
		organisationMember := &models.OrganisationMember{
			OrganisationID: org.ID,
			UserID:         member.ID,
			Role:           member.Role,
		}

		if err := organisationMember.SaveMember(tx); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, utils.FailedToCreateOrg)
			return
		}
	}

	// Try to commit to database
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"organisation": org})
}

func UpdateOrganisation(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteOrganisation(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

/**
 * Member related enpoints
 */
func GetOrganisationMembers(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func AddOrganisationMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func UpdateOrganisationMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteOrganisationMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
