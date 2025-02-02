package memberships

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
		ID       string `db:"id"`
		Username string `db:"username"`
		Name     string `db:"name"`
		Email    string `db:"email"`
		Password string `db:"password"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
		CreatedBy string `db:"created_by"`
		UpdatedBy string `db:"updated_by"`
	}
)
