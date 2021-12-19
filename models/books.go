package models

import (
	"database/sql"
	"fmt"
	"log"
)

// the json tags are to format the json response 
type Book struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	ImageURL    string `json:"imageURL"`
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

func (m BookModel) GetBooks() /* ([]Book,error) */ {
	//TODO
	/* rows, err := m.DB.Query("SELECT * FROM books;")
	if err != nil {
		log.Fatal(err)
	}

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		if err = rows.Scan(&bk.isbn, &bk.title, &bk.description, &bk.author, &bk.imageURL); err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	
	return bks,nil */
}

func (m BookModel) InsertBook(book Book){
	//TODO direct interaction with db layer done here
}

func (m BookModel) UpdateBook(book Book){
	//TODO direct interaction with db layer done here
}

func (m BookModel) DeleteBook(book Book){
	//TODO direct interaction with db layer done here
}

