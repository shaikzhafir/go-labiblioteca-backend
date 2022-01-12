package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-labiblioteca-backend/domain"
	"testing"
)

var book = &domain.Book{
	Isbn:        `lalala`,
	Title:       `asdasd`,
	Description: `asdasdwe`,
	Author:      `asdasd`,
	ImageURL:    `asdasdasd`,
}

func TestBookRepository_InsertBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO books(isbn, title, description, author, imageURL) VALUES($1, $2, $3, $4, $5);"

	mock.ExpectExec(query).WithArgs(book.Isbn, book.Title, book.Description, book.Author, book.ImageURL).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := BookRepository{db}
	_, err = repo.InsertBook(book)
	print(err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, err)
}
