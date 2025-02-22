package posts

import (
	"context"
	"database/sql"

	"github.com/chloeder/forum_app/internal/models/posts"
)

func (r *repository) GetPosts(ctx context.Context, limit, offset int) ([]*posts.PostModel, error) {
	query := `SELECT id, user_id, post_title, post_content, post_hastags, created_at, updated_at, created_by, updated_by FROM posts LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []*posts.PostModel
	for rows.Next() {
		var response posts.PostModel
		err := rows.Scan(&response.ID, &response.UserID, &response.PostTitle, &response.PostContent, &response.PostHastags, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
		if err != nil {
			return nil, err
		}
		responses = append(responses, &response)
	}

	return responses, nil
}

func (r *repository) GetPostById(ctx context.Context, id int64) (*posts.PostModel, error) {
	query := `SELECT id, post_title, post_content, post_hastags, created_at, updated_at, created_by, updated_by FROM posts WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)´
	var response posts.PostModel
	err := row.Scan(&response.ID, &response.PostTitle, &response.PostContent, &response.PostHastags, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &response, nil
}

func (r *repository) CreatePost(ctx context.Context, post *posts.PostModel) error {
	query := `INSERT INTO posts (id, user_id, post_title, post_content, post_hastags, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, post.ID, post.UserID, post.PostTitle, post.PostContent, post.PostHastags, post.CreatedAt, post.UpdatedAt, post.CreatedBy, post.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdatePost(ctx context.Context, id int64, post *posts.PostModel) error {
	query := `UPDATE posts SET post_title = ?, post_content = ?, post_hastags = ?, updated_at = ?, updated_by = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, post.PostTitle, post.PostContent, post.PostHastags, post.UpdatedAt, post.UpdatedBy, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeletePost(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
