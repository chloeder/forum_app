package memberships

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/memberships"
	"github.com/gin-gonic/gin"
)



type membershipService interface {
	SignUp(ctx context.Context, req *memberships.SignUpRequest) error
	SignIn(ctx context.Context, req *memberships.SignInRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipService membershipService
}

func NewHandler(api *gin.Engine, membershipService membershipService) *Handler {
	return &Handler{
		Engine: api,
		membershipService: membershipService,
	}
}

// AuthenticationRoute registers all routes for authentication
func (h *Handler) AuthenticationRoute() {
	router := h.Group("/memberships")

	router.GET("/register", h.Register)
	router.POST("/signup", h.SignUp)
	router.POST("/signin", h.SignIn)
}
