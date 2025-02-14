package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/chloeder/forum_app/internal/configs"
	"github.com/chloeder/forum_app/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.GetConfig().Service.SecretJWT

	return func(c *gin.Context) {
		// Get the token from the header
		token := c.GetHeader("Authorization")

		// Check if the token is empty
		header := strings.TrimSpace(token)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing authorization header"))
		}

		// Check if the token is valid
		user_id, username, email, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		}

		// Set the user ID, username, and email in the context
		c.Set("user_id", user_id)
		c.Set("username", username)
		c.Set("email", email)
		c.Next()
	}
}
