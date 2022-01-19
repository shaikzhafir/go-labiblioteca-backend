//this file serves as the connection layer between db and the app

package repository

import (
	"database/sql"
	"fmt"
	"go-labiblioteca-backend/domain"
	"log"
	"strings"
)

//this struct is connection to the db
type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
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

func (repository BookRepository) CreateBookTable() {
	qry := fmt.Sprintf(createTable, tableName)
	if _, err := repository.DB.Exec(qry); err != nil {
		log.Fatalf("error creating table %s", err)
	}
	log.Printf("database %s created/alr created", tableName)
}

func (repository BookRepository) GetBooks() ([]domain.Book, error) {
	rows, err := repository.DB.Query("SELECT * FROM books;")
	if err != nil {
		return nil, err
	}

	//bks is an array of Book
	bks := make([]domain.Book, 0)
	for rows.Next() {
		bk := domain.Book{}
		if err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Description, &bk.Author, &bk.ImageURL); err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
		fmt.Println(bks)
	}

	return bks, nil

}

func (repository BookRepository) InsertBook(book *domain.Book) (int64, error) {
	//TODO direct interaction with db layer done here
	result, err := repository.DB.Exec("INSERT INTO books (isbn, title, description, author, imageURL) VALUES ($1, $2, $3, $4, $5);", book.Isbn, book.Title, book.Description, book.Author, book.ImageURL)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil

}

func (repository BookRepository) UpdateBook(book domain.Book, isbn string) (int64, error) {
	//TODO direct interaction with db layer done here
	var arg []interface{}
	qry, args, err := updateStatement(isbn, arg, book)
	if err != nil {
		return 0, err
	}

	result, err := repository.DB.Exec(qry, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func updateStatement(isbn string, argArray []interface{}, updateData domain.Book) (string, []interface{}, error) {
	var queryString string
	i := 2
	argArray = append(argArray, isbn)
	if len(updateData.Title) > 0 {
		queryString += fmt.Sprintf(`title=$%d,`, i)
		argArray = append(argArray, updateData.Title)
		i++
	}
	if len(updateData.Description) > 0 {
		queryString += fmt.Sprintf(`description=$%d,`, i)
		argArray = append(argArray, updateData.Description)
		i++
	}
	if len(updateData.Author) > 0 {
		queryString += fmt.Sprintf(`author=$%d,`, i)
		argArray = append(argArray, updateData.Author)
		i++
	}
	if len(updateData.ImageURL) > 0 {
		queryString += fmt.Sprintf(`imageURL=$%d,`, i)
		argArray = append(argArray, updateData.ImageURL)
		i++
	}

	qry := fmt.Sprintf(`UPDATE books SET %s WHERE isbn=$1`, strings.TrimSuffix(queryString, ","))
	fmt.Println(qry)
	fmt.Printf(`%v+`, argArray)
	return qry, argArray, nil
}

func (repository BookRepository) DeleteBook(isbn string) (int64, error) {
	//TODO direct interaction with db layer done here
	result, err := repository.DB.Exec("DELETE FROM books WHERE isbn = $1", isbn)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
