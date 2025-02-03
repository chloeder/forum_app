package memberships

import "time"

type (
	SignUpRequest struct {
		Username string `json:"username"`
		Name 	 	 string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	UserModel struct {
		ID       int `db:"id"`
		Username string `db:"username"`
		Name     string `db:"name"`
		Email    string `db:"email"`
		Password string `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)
