package memberships

import (
	"net/http"

	"github.com/chloeder/forum_app/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn (c *gin.Context){
	ctx := c.Request.Context()

	var req memberships.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.membershipService.SignIn(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := memberships.SignInResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": response})
}
