package Database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func init() {
	var err error
	// Open the SQLite3 database file (creates it if it doesn't exist)
	db, err = sql.Open("sqlite3", "Database/ForumDatabase.db")
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}

	// Read the SQL schema from the Schema.sql file
	schemaFilePath := "Database/Schema.sql"
	schemaSQL, err := os.ReadFile(schemaFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read schema file: %v", err))
	}

	// Execute the SQL schema
	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		panic(fmt.Sprintf("Failed to execute schema: %v", err))
	}

	fmt.Println("Database initialized successfully!")
}
