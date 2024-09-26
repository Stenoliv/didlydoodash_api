package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrganisations(c *gin.Context) {
	organisations := []map[string]string{
		{"name": "Test"},
		{"name": "Test2"},
	}

	c.JSON(http.StatusOK, gin.H{"organisations": organisations})
}

func CreateOrganisation(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
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
