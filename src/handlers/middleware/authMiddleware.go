package middleware

import (
	"DidlyDoodash-api/src/data"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Auth check implementation
		data.CurrentUser = nil

		c.Next()
	}
}
