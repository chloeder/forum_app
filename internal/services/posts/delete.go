package posts

import (
	"context"
	"errors"
)

func (s *service) DeletePost(ctx context.Context, id int64) error {
	// Check if post exists
	_, err := s.postRepo.GetPostById(ctx, id)
	if err != nil {
		return err
	}

	// Delete post by ID
	err = s.postRepo.DeletePost(ctx, id)
	if err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}
