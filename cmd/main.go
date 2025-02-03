// Package main is the entry point of the application
package main

// Import necessary packages and dependencies
import (
	"log"

	"github.com/chloeder/forum_app/internal/configs"
	"github.com/chloeder/forum_app/internal/handlers/memberships"
	membershipsRepo "github.com/chloeder/forum_app/internal/repositories/memberships"
	membershipsService "github.com/chloeder/forum_app/internal/services/memberships"
	"github.com/chloeder/forum_app/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

// main function serves as the entry point of the application
func main() {
	// Initialize a new Gin router with default middleware
	r := gin.Default()

	// Declare config variable to store application configuration
	var (
		cfg *configs.Config
	)

	// Initialize configuration with specified options:
	// - Set config folder path
	// - Set config file name
	// - Set config file type
	if err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	); err != nil {
		// If configuration initialization fails, log the error and exit
		log.Fatalf("failed to initialize configs: %v", err)
	}

	// Get the configuration instance
	cfg = configs.GetConfig()
	// Log the configuration for debugging purposes
	log.Printf("config: %v", cfg)

	// Initialize the database connection
	db, err := internalsql.Connect(cfg.Database.DataSourcesName)
	if err != nil {
		// If database connection fails, log the error and exit
		log.Fatalf("failed to connect to the database: %v", err)
	}

	membershipRepo := membershipsRepo.NewRepository(db)
	membershipService := membershipsService.NewService(membershipRepo)

	// Initialize membership handler with the router
	membershipHandler := memberships.NewHandler(r, membershipService)
	// Register membership routes
	membershipHandler.RegisterRoute()

	// Start the server on the configured port
	r.Run(cfg.Service.Port)
}
