//the idea of this package is to create connection to DB, reuse the same instance in main.go

package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./biblioteca.db")
	if err != nil {
		log.Fatal(err)
	}
	// create the table if not exists
	return db, nil
}
