package posts

import (
	"context"
	"errors"
	"log"
	"strconv"
	time2 "time"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (s *service) LikedPost(ctx context.Context, postID int64, userID int64, req *posts.UserActivityRequest) error {
	time := time2.Now()

	userActivity := &posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   req.IsLiked,
		CreatedAt: time,
		UpdatedAt: time,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.LikedPost(ctx, userActivity)
	if err != nil {
		log.Println(err)
		return errors.New("failed to liked post")
	}

	return nil
}
