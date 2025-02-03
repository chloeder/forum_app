package memberships

import (
	"net/http"

	"github.com/chloeder/forum_app/internal/models/memberships"
	"github.com/gin-gonic/gin"
)


func (h *Handler) SignUp (c *gin.Context){
	ctx := c.Request.Context()

	var req memberships.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.membershipService.SignUp(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": req})
}
