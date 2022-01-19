package repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-labiblioteca-backend/domain"
	"regexp"
	"testing"
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
	print(err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, err)
}

func TestBookRepository_InsertBook_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO books (isbn, title, description, author, imageURL) VALUES ($1, $2, $3, $4, $5);"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(book.Isbn, book.Title, book.Description, book.Author, book.ImageURL).WillReturnError(fmt.Errorf("error"))

	repo := BookRepository{db}
	_, err = repo.InsertBook(book)
	fmt.Println(err)
	print(err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Error(t, err)
}
