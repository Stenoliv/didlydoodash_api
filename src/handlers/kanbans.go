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

	// Start transaction
	tx := db.DB.Begin()

	// Save kanban to transaction
	if err := kanban.SaveKanban(tx); err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.KanbanCreateError)
		return
	}

	// create default kanban category
	cateogry := &models.KanbanCategory{
		KanbanID: kanban.ID,
		Name:     "Not assigned",
	}

	// Save kanban category to transaction
	if err := cateogry.SaveCategory(tx); err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.KanbanCreateError)
		return
	}

	// Try to commit transaction
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.KanbanCreateError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"kanban": kanban})
}

func GetArchive(c *gin.Context) {
	kanbanID := c.Param("kanbanID")

	// Get archive from database
	archive, err := daos.GetKanbanArchive(kanbanID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.KanbanNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"archive": archive})
}
