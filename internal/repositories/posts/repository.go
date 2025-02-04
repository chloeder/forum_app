package posts

import "database/sql"

type repository struct {
	db *sql.DB
}

// NewRepository creates new repository instance with db connection
func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}
