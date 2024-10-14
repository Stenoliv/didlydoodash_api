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

func GetAnnouncements(c *gin.Context) {
	orgID := c.Param("id")
	announcements, err := daos.GetAnnouncements(orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	c.JSON(http.StatusOK, gin.H{"announcements": announcements})
}

type createAnnouncementInput struct {
	Text string `json:"text" binding:"required"`
}

func CreateAnnouncement(c *gin.Context) {
	var input createAnnouncementInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		fmt.Println("error in bindoby")
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	orgID := c.Param("id")
	tx := db.DB.Begin()
	organisation, err := daos.GetOrg(orgID)
	if err != nil || organisation == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ProjectNotFound)
		return
	}
	announcement := &models.Announcement{AnnouncmentText: input.Text, OrganisationID: orgID}
	if err = announcement.SaveAnnouncement(tx); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, utils.WhiteboardCreateError)
		tx.Rollback()
		return
	}
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.WhiteboardCreateError)
		tx.Rollback()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"announcement": announcement})
}
func DeleteAnnouncement(c *gin.Context) {
	id := c.Param("announcementID")

	announcement, err := daos.GetAnnouncement(id)
	if err != nil || announcement == nil {
		c.JSON(http.StatusBadRequest, utils.WhiteboardNotFound)
		return
	}
	// Try to delete organisation from database
	if err := db.DB.Delete(&announcement, "id = ?", announcement.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": announcement})
}
