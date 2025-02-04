package posts

import (
	"context"
	"errors"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (s *service) UpdatePost(ctx context.Context, post *posts.PostModel) error {
	// Check if post exists
	_, err := s.postRepo.GetPostById(ctx, post.ID)
	if err != nil {
		return err
	}

	// Update post by ID
	err = s.postRepo.UpdatePost(ctx, post)
	if err != nil {
		return errors.New("failed to update post")
	}

	return nil
}
