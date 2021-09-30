package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func ConnectDatabase() *sql.DB {
	var err error
	var faggot string = "midget"
	//sql.Open creates a pointer to Db (*Db)
	DbConn, err := sql.Open("postgres", "postgres://postgres:poop@localhost/bookdb?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(faggot)
	err = DbConn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return DbConn

}
