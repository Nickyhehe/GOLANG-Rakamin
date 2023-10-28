package database

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) *sql.DB {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func MigrateDB() {

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

}

func RunMigrations() {
	fmt.Println("Running database migrations...")
	MigrateDB()
	fmt.Println("Database migrations completed.")
}
