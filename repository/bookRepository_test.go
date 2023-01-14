package repository

import (
	"fmt"
	"go-labiblioteca-backend/domain"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var book = &domain.Book{
	Isbn:        `lalala`,
	Title:       `asdasd`,
	Description: `asdasdwe`,
	Author:      `asdasd`,
	ImageURL:    `asdasdasd`,
}

func TestBookRepository_InsertBook_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO books (isbn, title, description, author, imageURL) VALUES ($1, $2, $3, $4, $5);"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(book.Isbn, book.Title, book.Description, book.Author, book.ImageURL).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := BookRepository{db}
	_, err = repo.InsertBook(book)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, err)
}

var query = "INSERT INTO books (isbn, title, description, author, imageURL) VALUES ($1, $2, $3, $4, $5);"

func TestBookRepository_InsertBook_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(book.Isbn, book.Title, book.Description, book.Author, book.ImageURL).WillReturnError(fmt.Errorf("error"))

	repo := BookRepository{db}
	_, err = repo.InsertBook(book)

	assert.EqualError(t, err, "error")
}

func TestBookRepository_InsertManyBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(book.Isbn, book.Title, book.Description, book.Author, book.ImageURL).WillReturnResult(sqlmock.NewResult(1, 1))
	defer db.Close()

	// execute our actual logic
	repo := BookRepository{db}
	bookArray := []domain.Book{
		*book,
	}
	err = repo.InsertManyBooks(&bookArray)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestBookRepository_DeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM books WHERE isbn = $1"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(book.Isbn).WillReturnResult(sqlmock.NewResult(1, 1))
	defer db.Close()

	// execute our actual logic
	repo := BookRepository{db}
	deleteBook, err := repo.DeleteBook(book.Isbn)
	assert.NoError(t, err)
	assert.Equal(t, deleteBook, int64(1))
}

func TestBookRepository_UpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := "UPDATE books SET title=$2,description=$3,author=$4,imageURL=$5 WHERE isbn=$1"
	t.Logf("the query is %s\n", query)
	mock.ExpectExec(query).WithArgs(book.Isbn, book.Title, book.Description, book.Author, book.ImageURL).WillReturnResult(sqlmock.NewResult(1, 1))
	defer db.Close()

	// execute our actual logic
	repo := BookRepository{db}
	updateBook, err := repo.UpdateBook(*book, book.Isbn)
	assert.NoError(t, err)
	assert.Equal(t, updateBook, int64(1))
}

func TestBookRepository_InsertBook(t *testing.T) {

}
