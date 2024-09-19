package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	main := r.Group("/")
	{
		main.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Welocme to didlydoodash api"})
		})
	}

	r.Run("0.0.0.0:3000")
}
