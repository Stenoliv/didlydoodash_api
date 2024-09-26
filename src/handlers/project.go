package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

func CreateProjects(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
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
