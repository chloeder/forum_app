package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/chloeder/forum_app/internal/models/posts"
)



func (s *service) GetCommentsByPostId(ctx context.Context, postID int64, limit, offset int) ([]*posts.CommentModel, error) {
	post, err := s.postRepo.GetCommentsByPostId(ctx, postID, limit, offset)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *service) CreateComment(ctx context.Context, userID int64, req *posts.CreateCommentRequest) error {
	time := time.Now()

	comment := &posts.CommentModel{
		PostID: req.PostID,
		Comment: req.Comment,
		CreatedAt: time,
		UpdatedAt: time,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreateComment(ctx, comment)
	if err != nil {
		return errors.New("failed to create comment")
	}

	return nil
}
