package Database

import (
	"database/sql"
	"fmt"
	"os"
)

var Db *sql.DB

func init() {
	var err error

	// Check if the database file exists
	dbFilePath := "ForumDatabase.db"
	dbFileExists := fileExists(dbFilePath)

	// Open the SQLite3 database file (creates it if it doesn't exist)
	Db, err = sql.Open("sqlite3", dbFilePath)
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
	_, err = Db.Exec(string(schemaSQL))
	if err != nil {
		panic(fmt.Sprintf("Failed to execute schema: %v", err))
	}

	if !dbFileExists {
		// Read the SQL data from the Data.sql file if the database file didn't exist
		dataFilePath := "Database/Data.sql"
		dataSQL, err := os.ReadFile(dataFilePath)
		if err != nil {
			panic(fmt.Sprintf("Failed to read data file: %v", err))
		}

		// Execute the SQL data
		_, err = Db.Exec(string(dataSQL))
		if err != nil {
			panic(fmt.Sprintf("Failed to execute data: %v", err))
		}
	}

	fmt.Println("Database initialized successfully!")
}

// Helper function to check if a file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
