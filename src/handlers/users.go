package handlers

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users paginated
func GetAllUsers(c *gin.Context) {
	result, err := daos.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
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
	usr, err := daos.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	if (usr == models.User{}) { // Zero-value check
		c.JSON(http.StatusNotFound, utils.UserNotFound)
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
