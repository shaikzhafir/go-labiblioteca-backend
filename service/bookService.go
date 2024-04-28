package service

import (
	"context"
	_ "fmt"
	"labiblioteca/sqlcgen"
)

type BookService struct {
	sqlcgen sqlcgen.Queries
}

// this is used to init in main.go
func NewBookService(sqlcgen sqlcgen.Queries) *BookService {
	return &BookService{sqlcgen: sqlcgen}
}

func (service *BookService) GetBooks(ctx context.Context) ([]sqlcgen.Book, error) {
	return service.sqlcgen.ListBooks(ctx)
}

func (service *BookService) AddBook(ctx context.Context, insertParams *sqlcgen.InsertBookParams) error {
	return service.sqlcgen.InsertBook(ctx, *insertParams)
}

func (service *BookService) DeleteBook(ctx context.Context, id int64) error {
	return service.sqlcgen.DeleteBook(ctx, id)
}

func (service *BookService) UpdateBook(ctx context.Context, updateParams *sqlcgen.UpdateBookParams) error {
	return service.sqlcgen.UpdateBook(ctx, *updateParams)
}

func (service *BookService) GetBooksByAuthor(ctx context.Context, author string) ([]sqlcgen.Book, error) {
	return service.sqlcgen.GetBooksByAuthor(ctx, author)
}

func (service *BookService) GetBookByID(ctx context.Context, id int64) (sqlcgen.Book, error) {
	return service.sqlcgen.GetBookByID(ctx, id)
}
