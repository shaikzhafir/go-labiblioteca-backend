//the idea of this package is to create connection to DB, reuse the same instance in main.go

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	var err error
	getEnv()
	//sql.Open creates a pointer to Db (*Db)
	//DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB")))

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return db, nil

}

func getEnv() {
	fmt.Printf("password is %s", os.Getenv("POSTGRES_PASSWORD"))
}
