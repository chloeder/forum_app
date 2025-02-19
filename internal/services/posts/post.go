package posts

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (s *service) GetPosts (ctx context.Context, limit, offset int) ([]*posts.PostModel, error) {
	posts, err := s.postRepo.GetPosts(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

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

func (s *service) UpdatePost(ctx context.Context,	id int64, post *posts.UpdatePostRequest) error {
	// Check if post exists
	existingPost, err := s.postRepo.GetPostById(ctx, id)
	if err != nil {
		return err
	}

	// Join post hastags
	postHastags := strings.Join(post.PostHastags, ",")
	time := time.Now()

   // Apply patches only if fields are provided
	if post.PostTitle != nil {
			existingPost.PostTitle = *post.PostTitle
	}
	if post.PostContent != nil {
			existingPost.PostContent = *post.PostContent
	}
	if post.PostHastags != nil {
			existingPost.PostHastags = postHastags
	}

	// Update post
	updatePost := &posts.PostModel{
		ID: id,
		PostTitle: existingPost.PostTitle,
		PostContent: existingPost.PostContent,
		PostHastags: existingPost.PostHastags,
		CreatedAt: time,
		UpdatedAt: time,
		CreatedBy: strconv.FormatInt(id, 10),
		UpdatedBy: strconv.FormatInt(id, 10),
	}

	err = s.postRepo.UpdatePost(ctx, id, updatePost)
	if err != nil {
		return errors.New("failed to update post")
	}

	return nil
}

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
