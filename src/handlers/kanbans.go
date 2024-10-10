package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllKanbans(c *gin.Context) {
	projectID := c.Param("projectID")
	kanbans, err := daos.GetAllKanbans(projectID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.KanbanNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"kanbans": kanbans})
}

// Create new kanban
type createKanbanInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateKanban(c *gin.Context) {
	var input createKanbanInput
	projectID := c.Param("projectID")
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	kanban := &models.Kanban{
		Name:      input.Name,
		ProjectID: projectID,
	}

	tx := db.DB.Begin()

	if err := kanban.SaveKanban(tx); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.KanbanCreateError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.KanbanCreateError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"kanban": kanban})
}
