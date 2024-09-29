package handlers

import (
	"DidlyDoodash-api/src/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users paginated
func GetAllUsers(c *gin.Context) {
	result, err := models.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get a single user
func GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		id = *models.CurrentUser
	}
	usr, err := models.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if (usr == models.User{}) { // Zero-value check
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found", "errorcode": "404"})
		return
	}

	c.JSON(http.StatusOK, usr)
}

// Put profile/user
func PutUser(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

// Patch profile/user
func PatchUser(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}
