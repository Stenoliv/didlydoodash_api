package main

import (
	"DidlyDoodash-api/src/config"
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/handlers"
	"DidlyDoodash-api/src/handlers/middleware"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.RequestLoggerMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowMethods:     []string{"POST", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init connection to database
	db.Init()

	// Auth endpoints
	auth := r.Group("/auth")
	{
		auth.POST("/signin", handlers.Signin)   // Login user
		auth.POST("/signup", handlers.Signup)   // Register new user
		auth.POST("/signout", handlers.Signout) // Logout user
		auth.GET("/refresh", handlers.Refresh)  // Refresh access token
	}

	// Users endpoints
	users := r.Group("/users", middleware.AuthMiddleware())
	{
		users.GET("", handlers.GetAllUsers)
		users.GET("/:id", handlers.GetUser)
		users.PUT("/:id", handlers.PutUser)
		users.PATCH("/:id", handlers.PatchUser)
	}

	// Organisations endpoints
	organisation := r.Group("/organisations", middleware.AuthMiddleware())
	{
		// Basic endpoints
		organisation.GET("", handlers.GetOrganisations)          // Get organisation user is part of
		organisation.POST("", handlers.CreateOrganisation)       // Create a new organisation
		organisation.PATCH("/:id", handlers.UpdateOrganisation)  // Update organisation
		organisation.DELETE("/:id", handlers.DeleteOrganisation) // Delete organisation

		// Organisation members
		organisation.GET("/:id/members", handlers.GetOrganisationMembers)              // Get organisation members
		organisation.POST("/:id/members", handlers.AddOrganisationMember)              // Add member to organisation
		organisation.PATCH("/:id/members/:userID", handlers.UpdateOrganisationMember)  // Update role etc... of organisation member
		organisation.DELETE("/:id/members/:userID", handlers.DeleteOrganisationMember) // Remove organisation member

		organisation.GET("/:id/chats", handlers.GetChats)
		organisation.POST("/:id/chats", handlers.CreateChat)
	}

	// Projects endpoints
	project := r.Group("/project", middleware.AuthMiddleware())
	{
		// Basic endpoints
		project.GET("/:id", handlers.GetProjects)                              // Get all projects of selected organisation
		project.POST("/:id", handlers.CreateProjects)                          // Create a new project in organisation
		project.PATCH("/:id/:projectID", handlers.UpdateProjects)              // Update a project in an organisation
		project.DELETE("/:id/:projectID", handlers.DeleteProjects)             // Archive a project in an organisation
		project.DELETE("/:id/:projectID/delete", handlers.PermaDeleteProjects) // Delete a project in an organisation

		// Project members
		project.GET("/:id/members", handlers.GetProjectMembers)              // Get project members
		project.POST("/:id/members", handlers.GetProjectMember)              // Add project member
		project.PATCH("/:id/members/:userID", handlers.UpdateProjectMember)  // Update project member
		project.DELETE("/:id/members/:userID", handlers.DeleteProjectMember) // Remove member from project
	}

	r.Run(fmt.Sprintf(":%s", config.PORT))
}
