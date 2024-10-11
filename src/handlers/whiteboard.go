package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWhiteboards(c *gin.Context) {
	projID := c.Param("id")
	whiteboards, err := daos.GetWhiteboards(projID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Whiteboards": whiteboards})
}
func CreateNewWhiteboard(c *gin.Context) {
	projID := c.Param("id")
	tx := db.DB.Begin()

}