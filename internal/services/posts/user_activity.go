package posts

import (
	"context"
	"errors"
	"strconv"
	time2 "time"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (s *service) LikedPost(ctx context.Context, postID int64, userID int64) error {
	time := time2.Now()

	userActivity := &posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		CreatedAt: time,
		UpdatedAt: time,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	if userActivity.IsLiked == true {
		userActivity.IsLiked = false
	} else {
		userActivity.IsLiked = true
	}

	err := s.postRepo.LikedPost(ctx, userActivity)
	if err != nil {
		return errors.New("failed to liked post")
	}

	return nil
}
