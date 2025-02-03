package memberships

import "time"

type (
	SignUpRequest struct {
		Username string `json:"username"`
		Name 	 	 string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	SignInResponse struct {
		AccessToken string `json:"access_token"`
	}
)

type (
	UserModel struct {
		ID       int64 `db:"id"`
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
