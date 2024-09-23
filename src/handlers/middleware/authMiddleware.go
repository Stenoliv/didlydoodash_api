package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		// Auth check implementation
		
		ctx.Next()
	}
}