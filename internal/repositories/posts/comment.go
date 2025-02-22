package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (r *repository) GetCommentsByPostId(ctx context.Context, postID int64, limit, offset int) ([]*posts.CommentModel, error) {
	query := `SELECT id, post_id, user_id, comment, created_at, updated_at, created_by, updated_by FROM comments WHERE post_id = ? LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []*posts.CommentModel
	for rows.Next() {
		var response posts.CommentModel
		err := rows.Scan(&response.ID, &response.PostID, &response.UserID, &response.Comment, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
		if err != nil {
			return nil, err
		}
		responses = append(responses, &response)
	}

	return responses, nil
}

func (r *repository) CreateComment(ctx context.Context, comment *posts.CommentModel) error {
	query := `INSERT INTO comments (id, post_id, user_id, comment, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, comment.ID, comment.PostID, comment.UserID, comment.Comment, comment.CreatedAt, comment.UpdatedAt, comment.CreatedBy, comment.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
