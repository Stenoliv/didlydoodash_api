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