package main

import (
	"DidlyDoodash-api/src/config"
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/handlers"
	"DidlyDoodash-api/src/handlers/middleware"
	"DidlyDoodash-api/src/ws/chat"
	"DidlyDoodash-api/src/ws/kanban"
	"DidlyDoodash-api/src/ws/whiteboardws"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the Gin mode based on the environment variable
	if os.Getenv("MODE") == "production" {
		gin.SetMode(gin.ReleaseMode) // Default to production mode
	} else {
		gin.SetMode(gin.DebugMode) // Default to debug mode
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLoggerMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                 // Add your frontend URL
		AllowMethods:     []string{"POST", "PUT", "PATCH", "GET", "DELETE", "OPTIONS"},  // Include OPTIONS
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "_retry"}, // Specify headers
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
		chatHandler := chat.NewChatHandler()

		// Basic endpoints
		organisation.GET("", handlers.GetOrganisations)          // Get organisation user is part of
		organisation.POST("", handlers.CreateOrganisation)       // Create a new organisation
		organisation.PATCH("/:id", handlers.UpdateOrganisation)  // Update organisation
		organisation.DELETE("/:id", handlers.DeleteOrganisation) // Delete organisation

		// Organisation members
		organisation.GET("/:id/members", handlers.GetOrganisationMembers)              // Get organisation members
		organisation.POST("/:id/members/:userID", handlers.AddOrganisationMember)      // Add member to organisation
		organisation.PATCH("/:id/members/:userID", handlers.UpdateOrganisationMember)  // Update role etc... of organisation member
		organisation.DELETE("/:id/members/:userID", handlers.DeleteOrganisationMember) // Remove organisation member

		// Organisation chats
		organisation.GET("/:id/chats", handlers.GetChats)
		organisation.POST("/:id/chats", handlers.CreateChat)
		organisation.PUT("/:id/chats/:chatId/member/:userId", handlers.AddUserToChat)
		organisation.DELETE("/:id/chats/:chatId/member/:userId", handlers.RemoveUserFromChat)
		organisation.GET("/:id/chats/:chatId", chatHandler.JoinRoom)
		//announcements
		organisation.GET("/:id/announcements", handlers.GetAnnouncements)
		organisation.DELETE("/:id/announcements/:announcementID", handlers.DeleteAnnouncement)
		organisation.POST("/:id/announcements", handlers.CreateAnnouncement)
		// Organisation chats notifcations
		organisation.GET("/:id/chats/notifications", chatHandler.NotificationHandler)

		// Projects endpoints
		project := organisation.Group("/:id/projects", middleware.ProjectMiddleware())
		{
			// Basic endpoints
			project.GET("", handlers.GetAllProjects)                           // Get all projects of selected organisation
			project.POST("", handlers.CreateProjects)                          // Create a new project in organisation
			project.PATCH("/:projectID", handlers.UpdateProjects)              // Update a project in an organisation
			project.DELETE("/:projectID", handlers.DeleteProjects)             // Archive a project in an organisation
			project.DELETE("/:projectID/delete", handlers.PermaDeleteProjects) // Delete a project in an organisation

			// Project members
			project.GET("/:projectID/members", handlers.GetProjectMembers)              // Get project members
			project.POST("/:projectID/members", handlers.GetProjectMember)              // Add project member
			project.PATCH("/:projectID/members/:userID", handlers.UpdateProjectMember)  // Update project member
			project.DELETE("/:projectID/members/:userID", handlers.DeleteProjectMember) // Remove member from project

			// Kanban endpoints
			kanbans := project.Group("/:projectID/kanbans")
			{
				kanbanHandler := kanban.NewHandler()

				// Basic endpoints
				kanbans.GET("", handlers.GetAllKanbans)
				kanbans.POST("", handlers.CreateKanban)
				kanbans.DELETE("/:kanbanID", nil)

				// Kanban archive
				kanbans.GET("/:kanbanID/archive", handlers.GetArchive)

				// WS
				kanbans.GET("/:kanbanID", kanbanHandler.JoinKanban)
			}

			// Whiteboard
			whiteboard := project.Group("/:projectID/whiteboards")
			{
				// Basic endpoints
				whiteboard.GET("", handlers.GetWhiteboards)                    // Get whiteboard user is part of
				whiteboard.POST("", handlers.CreateNewWhiteboard)              // Create a new Whiteboard
				whiteboard.DELETE("/:whiteboardID", handlers.DeleteWhiteboard) // Delete Whiteboard
				whiteboardHandler := whiteboardws.NewHandler()
				whiteboard.GET("/:wbID", whiteboardHandler.HandleConnections)
			}
		}

	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.PORT),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
