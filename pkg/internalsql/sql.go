package internalsql

import (
	"database/sql"
	"log"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	// Connect to the database using the specified data source name
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		// If an error occurs while connecting to the database, log the error
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Return the database connection
	return db, nil
}
