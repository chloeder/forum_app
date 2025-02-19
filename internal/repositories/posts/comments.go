package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/comments"
)

func (r *repository) GetCommentsByPostId(ctx context.Context, limit, offset int) ([]*comments.CommentModel, error) {
	query := `SELECT id, post_id, comment, created_at, updated_at, created_by, updated_by FROM comments WHERE post_id = ? LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []*comments.CommentModel
	for rows.Next() {
		var response comments.CommentModel
		err := rows.Scan(&response.ID, &response.PostID, &response.Comment, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
		if err != nil {
			return nil, err
		}
		responses = append(responses, &response)
	}

	return responses, nil
}

func (r *repository) CreateComment(ctx context.Context, comment *comments.CommentModel) error {
	query := `INSERT INTO comments (id, post_id, comment, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, comment.ID, comment.PostID, comment.Comment, comment.CreatedAt, comment.UpdatedAt, comment.CreatedBy, comment.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
