package posts

import (
	"context"
	"errors"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (s *service) GetPostById(ctx context.Context, id int64) (*posts.PostModel, error) {
	post, err := s.postRepo.GetPostById(ctx, id)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, errors.New("post not found")
	}

	return post, nil
}
