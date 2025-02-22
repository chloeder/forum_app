package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/posts"
)

type postRepository interface {
	GetPosts(ctx context.Context, limit, offset int) ([]*posts.PostModel, error)
	GetPostById(ctx context.Context, id int64) (*posts.PostModel, error)
	CreatePost(ctx context.Context, post *posts.PostModel) error
	UpdatePost(ctx context.Context, id int64, post *posts.PostModel) error
	DeletePost(ctx context.Context, id int64) error
	CreateComment(ctx context.Context, comment *posts.CommentModel) error
	LikedPost(ctx context.Context, userActivity *posts.UserActivityModel) error
}

type service struct {
	postRepo postRepository
}

// NewService creates new service instance with repository
func NewService(postRepo postRepository) *service {
	return &service{
		postRepo: postRepo,
	}
}
