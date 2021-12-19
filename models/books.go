package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Book struct {
	Isbn        string
	Title       string
	Description string
	Author      string
	ImageURL    string
}

//this struct is connection to the db
type BookModel struct {
	DB *sql.DB
}

const (
	tableName   = "books"
	createTable = `CREATE TABLE IF NOT EXISTS %s (
		isbn TEXT PRIMARY KEY,
		title TEXT,
		description TEXT,
		author TEXT,
		imageUrl TEXT
	)`
)

func (m BookModel) CreateBookTable() {
	qry := fmt.Sprintf(createTable, tableName)
	if _, err := m.DB.Exec(qry); err != nil {
		log.Fatalf("error creating table %s", err)
	}
	log.Printf("database %s created/alr created", tableName)
}


func (m BookModel) InsertBook(book Book){
	
}
