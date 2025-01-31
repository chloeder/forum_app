package main

import (
	"github.com/chloeder/forum_app/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

	// Create a new handler for memberships
	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()


	r.Run(":8080") // Listen and serve on localhost:8080
}
