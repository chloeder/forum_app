package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})
}
