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

func GetOrganisations(c *gin.Context) {
	orgs, err := daos.GetAllOrgs()
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
	User models.User `json:"user"`
	Role string      `json:"role"`
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
		Name: input.Name,
	}

	// Save organisation object to database
	if err := org.SaveOrganisation(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.FailedToCreateOrg)
		return
	}

	var owner *models.OrganisationMember

	// Create organisation members
	for _, member := range input.Members {

		organisationMember := &models.OrganisationMember{
			OrganisationID: org.ID,
			Organisation:   org,
			UserID:         member.User.ID,
			Role:           datatypes.ToOrganisationRole(member.Role),
		}

		if datatypes.CEO == datatypes.ToOrganisationRole(member.Role) {
			if owner == nil {
				owner = organisationMember
			} else {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, utils.InvalidInput)
				return
			}
		}

		if err := organisationMember.SaveMember(tx); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, utils.FailedToCreateOrg)
			return
		}

		org.Members = append(org.Members, *organisationMember)
	}

	// Set owner of organisation
	org.Owner = *owner

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

type deleteOrganisationInput struct {
	Name     string `json:"orgName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func DeleteOrganisation(c *gin.Context) {
	id := c.Param("id")
	var input deleteOrganisationInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Check that organisation exists
	org, err := daos.GetOrg(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Get user from JWT
	user, err := daos.GetUser(*models.CurrentUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Check if user is owner
	if org.Owner.User.ID != user.ID {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Check user password
	if !user.Validatepassword(input.Password) {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Check that input name matches organisation name
	if org.Name != input.Name {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Try to delete organisation from database
	if err := db.DB.Delete(&org).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": org})
}

/**
 * Member related enpoints
 */
func GetOrganisationMembers(c *gin.Context) {
	id := c.Param("id")
	members, err := daos.GetMembers(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	c.JSON(http.StatusOK, gin.H{"members": members, "organisationId": id})
}

func AddOrganisationMember(c *gin.Context) {
	id := c.Param("id")
	userId := c.Param("userID")
	role := datatypes.ToOrganisationRole(c.Query("role"))
	tx := db.DB.Begin()

	member := &models.OrganisationMember{
		OrganisationID: id,
		UserID:         userId,
		Role:           role,
	}

	if err := member.SaveMember(tx); err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"member": member})
}

func UpdateOrganisationMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteOrganisationMember(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
