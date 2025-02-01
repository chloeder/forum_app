package memberships

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

// NewRepository creates new repository instance with db connection
func NewRepository(db *sql.DB) *repository {
	// Query all users
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("failed to query users ", err)
	}

	defer rows.Close()

	// Read and log each user
	for rows.Next() {
		var id int
		var email string
		if err = rows.Scan(&id, &email); err != nil {
			log.Println("failed to scan row ", err)
		}
		log.Printf("id: %v, email: %v", id, email)
	}

	// Return repository with db
	return &repository{db}
}
