package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string `json:"post_title"`
		PostContent string `json:"post_content"`
		PostHastags []string `json:"post_hastags"`
	}

	UpdatePostRequest struct {
		PostTitle   *string `json:"post_title,omitempty"`
		PostContent *string `json:"post_content,omitempty"`
		PostHastags []string `json:"post_hastags,omitempty"`
	}

		CreateCommentRequest struct {
		PostID  int64  `json:"post_id"`
		Comment string `json:"comment"`
	}
)

type (
	PostModel struct {
		ID          int64 `db:"id"`
		UserID      int64 `db:"user_id"`
		PostTitle   string `db:"post_title"`
		PostContent string `db:"post_content"`
		PostHastags string `db:"post_hastags"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		CreatedBy   string `db:"created_by"`
		UpdatedBy   string `db:"updated_by"`
	}

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
