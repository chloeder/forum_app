package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (s *service) GetPosts (ctx context.Context, limit, offset int) ([]*posts.PostModel, error) {
	posts, err := s.postRepo.GetPosts(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
