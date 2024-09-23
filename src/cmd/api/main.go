package main

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/handlers"
	"DidlyDoodash-api/src/handlers/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.RequestLoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())

	// Init connection to database
	db.Init()

	auth := r.Group("/")
	{
		auth.POST("/signin", handlers.Signin)
		auth.POST("/signup", handlers.Signup)
	}

	organisation := r.Group("/organisations", middleware.AuthMiddleware())
	{
		organisation.GET("", handlers.GetOrganisations)
	}

	r.Run("0.0.0.0:3000")
}
