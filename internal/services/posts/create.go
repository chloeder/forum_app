package posts

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/chloeder/forum_app/internal/models/posts"
)


func (s *service) CreatePost(ctx context.Context, userID int64, req *posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHastags, ",")
	time := time.Now()

	post := &posts.PostModel{
		UserID: userID,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHastags,
		CreatedAt: time,
		UpdatedAt: time,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreatePost(ctx, post)
	if err != nil {
		return errors.New("failed to create post")
	}

	return nil
}
