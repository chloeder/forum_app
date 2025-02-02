package memberships

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

// NewRepository creates new repository instance with db connection
func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}
