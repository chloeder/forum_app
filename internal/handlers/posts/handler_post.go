package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/middleware"
	"github.com/chloeder/forum_app/internal/models/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	GetPosts(ctx context.Context, limit, offset int) ([]*posts.PostModel, error)
	GetPostById(ctx context.Context, id int64) (*posts.PostModel, error)
	CreatePost(ctx context.Context, userID int64, req *posts.CreatePostRequest) error
	UpdatePost(ctx context.Context, id int64, req *posts.UpdatePostRequest) error
	DeletePost(ctx context.Context, id int64) error

	GetCommentsByPostId(ctx context.Context, postID int64, limit, offset int) ([]*posts.CommentModel, error)
	CreateComment(ctx context.Context, userID int64, req *posts.CreateCommentRequest) error
}

type Handler struct {
	*gin.Engine

	postService postService
}

// NewHandler creates new handler instance with service
func NewHandler(api *gin.Engine, postService postService) *Handler {
	return &Handler{
		Engine:      api,
		postService: postService,
	}
}

// PostRoute registers all routes for post handler
func (h *Handler) PostRoute(){
	router := h.Group("/posts")
	router.Use(middleware.AuthMiddleware())

	router.GET("/", h.GetPosts)
	router.GET("/:id", h.GetPostById)
	router.POST("/", h.CreatePost)
	router.PATCH("/:id", h.UpdatePost)
	router.DELETE("/:id", h.DeletePost)
}
