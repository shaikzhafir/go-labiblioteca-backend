package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	//sql.Open creates a pointer to Db (*Db)
	DbConn, err := sql.Open("postgres", "postgres://poop@localhost/bookdb?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	DbConn.SetMaxOpenConns(3)

}
