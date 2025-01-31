package memberships

import "github.com/gin-gonic/gin"

type Handler struct {
	*gin.Engine
}

func NewHandler(api *gin.Engine) *Handler {
	return &Handler{api}
}

func (h *Handler) RegisterRoute() {
	router := h.Group("/memberships")

	router.GET("/register", h.Register)
}
