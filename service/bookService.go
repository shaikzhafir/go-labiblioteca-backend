package service

import (
	_ "fmt"
	"go-labiblioteca-backend/domain"
)

type BookRepository interface {
	GetBooks() ([]domain.Book, error)
	InsertBook(*domain.Book) (int64, error)
	UpdateBook(domain.Book, string) (int64, error)
	DeleteBook(string) (int64, error)
}

type BookService struct {
	repository BookRepository
}

// this is used to init in main.go
func NewBookService(repository BookRepository) *BookService {
	return &BookService{repository: repository}
}

func (service *BookService) GetBooks() ([]domain.Book, error) {
	bks, err := service.repository.GetBooks()
	if err != nil {
		return nil, err
	}
	return bks, err
}

func (service *BookService) AddBook(book *domain.Book) (int64, error) {
	rowsAffected, err := service.repository.InsertBook(book)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (service *BookService) DeleteBook(isbn string) (int64, error) {
	rowsAffected, err := service.repository.DeleteBook(isbn)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (service *BookService) UpdateBook(book domain.Book, isbn string) (int64, error) {
	rowsAffected, err := service.repository.UpdateBook(book, isbn)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
