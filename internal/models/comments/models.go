package comments

import "time"

type (
	CreateCommentRequest struct {
		PostID  int64  `json:"post_id"`
		Comment string `json:"comment"`
	}
)

type (
	CommentModel struct {
		ID        int64     `db:"id"`
		PostID    int64     `db:"post_id"`
		Comment   string    `db:"comment"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}
)
