package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (r *repository) CreateComment(ctx context.Context, comment *posts.CommentModel) error {
	query := `INSERT INTO comments (id, post_id, user_id, comment, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, comment.ID, comment.PostID, comment.UserID, comment.Comment, comment.CreatedAt, comment.UpdatedAt, comment.CreatedBy, comment.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
