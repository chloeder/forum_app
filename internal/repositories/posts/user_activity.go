package posts

import (
	"context"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (r *repository) LikedPost(ctx context.Context, userActivity *posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (id, post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ? ,? ,? ,?)`

	_, err := r.db.ExecContext(ctx, query, userActivity.ID, userActivity.PostID, userActivity.UserID, userActivity.IsLiked, userActivity.CreatedAt, userActivity.UpdatedAt, userActivity.CreatedBy, userActivity.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
