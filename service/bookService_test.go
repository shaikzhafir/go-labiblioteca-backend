package service

import (
	"github.com/stretchr/testify/assert"
	"go-labiblioteca-backend/domain"
	mocks "go-labiblioteca-backend/mocks/service"
	"testing"
)

var book = &domain.Book{
	Isbn:        `lalala`,
	Title:       `asdasd`,
	Description: `asdasdwe`,
	Author:      `asdasd`,
	ImageURL:    `asdasdasd`,
}

func TestBookService_GetBooks(t *testing.T) {
	repo := mocks.BookRepository{}
	repo.On("GetBooks").Return([]domain.Book{*book}, nil)

	// execute the service
	svc := NewBookService(&repo)
	books, err := svc.GetBooks()
	assert.NoError(t, err)
	assert.Equal(t, books, []domain.Book{*book})
}
