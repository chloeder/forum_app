// Package main is the entry point of the application
package main

// Import necessary packages and dependencies
import (
	"log"

	"github.com/chloeder/forum_app/internal/configs"
	"github.com/chloeder/forum_app/internal/handlers/memberships"
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

	// Initialize membership handler with the router
	membershipHandler := memberships.NewHandler(r)
	// Register membership routes
	membershipHandler.RegisterRoute()

	// Start the server on the configured port
	r.Run(cfg.Service.Port)
}
