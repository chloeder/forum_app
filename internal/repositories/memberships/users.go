package memberships

import (
	"context"
	"database/sql"

	"github.com/chloeder/forum_app/internal/models/memberships"
)

func (r *repository) GetUser (ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, username, name, email, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = $1 OR username = $2`

	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Username, &response.Name, &response.Email, &response.Password, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser (ctx context.Context, user *memberships.UserModel) error {
	query := `INSERT INTO users (id, username, name, email, password, created_at, updated_at, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
