package main

import (
	"DidlyDoodash-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	main := r.Group("/")
	{
		main.POST("/signin", handlers.Signin)
		main.POST("/signup", handlers.Signup)
	}

	r.Run("0.0.0.0:3000")
}
