package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWhiteboards(c *gin.Context) {
	projID := c.Param("projectID")
	whiteboards, err := daos.GetWhiteboards(projID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	c.JSON(http.StatusOK, gin.H{"whiteboards": whiteboards})
}

type createWhiteboardInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateNewWhiteboard(c *gin.Context) {
	var input createWhiteboardInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	projID := c.Param("projectID")
	tx := db.DB.Begin()
	project, err := daos.GetProject(projID)
	if err != nil || project == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ProjectNotFound)
		return
	}
	userExists := false
	fmt.Println(len(project.Members), *models.CurrentUser)
	for _, user := range project.Members {
		if *models.CurrentUser == user.UserID {
			userExists = true
			break
		}
	}
	if !userExists {
		c.AbortWithStatusJSON(http.StatusForbidden, utils.WhiteboardCreateError)
		return
	}
	whiteboard := &models.WhiteboardRoom{ProjectID: projID, Name: input.Name}
	if err = whiteboard.SaveWhiteboardRoom(tx); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, utils.WhiteboardCreateError)
		tx.Rollback()
		return
	}
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.WhiteboardCreateError)
		tx.Rollback()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"whiteboard": whiteboard})
}

func DeleteWhiteboard(c *gin.Context) {
	id := c.Param("whiteboardID")

	whiteboard, err := daos.GetWhiteboard(id)
	if err != nil || whiteboard == nil {
		c.JSON(http.StatusBadRequest, utils.WhiteboardNotFound)
		return
	}
	// Try to delete organisation from database
	if err := db.DB.Delete(&whiteboard, "id = ?", whiteboard.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": whiteboard})
}
