package handlers

import (
	"encoding/json"
	"fmt"
	"go-labiblioteca-backend/domain"
	"go-labiblioteca-backend/service"
	"io/ioutil"
	"log"
	"net/http"
	_ "strings"
)

type BookStore struct {
	bookList []string
	authors  []string
}

type BookHandler struct {
	service *service.BookService
}

// NewBookHandler for it to be called by outer domain
func NewBookHandler(service *service.BookService) BookHandler {
	return BookHandler{service: service}
}

func (b *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		b.getBook(w)
	case http.MethodPost:
		b.postBook(w, r)
	case http.MethodPut:
		b.updateBook(w, r)
	case http.MethodDelete:
		b.deleteBook(w, r)
	}

}

func (b *BookHandler) getBook(w http.ResponseWriter) {
	bks, err := b.service.GetBooks()
	if err != nil {
		fmt.Println("haha")
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bks); err != nil {
		log.Println(err)
	}
}

func (b *BookHandler) postBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	//handle error for decoding
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = b.service.AddBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Book %s created successfully", book.Title)
}

func (b *BookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	//extract isbn first
	isbn := r.URL.Query().Get("isbn")
	book, err := unmarshalBook(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rowsAffected, err := b.service.UpdateBook(book, isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Book %s updated successfully (%d row affected)\n", isbn, rowsAffected)
	//
}

func unmarshalBook(r *http.Request) (book domain.Book, err error) {
	data, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(data, &book)
	return book, err
}

// a general parser that takes in variables of type interface{}
func dataParser(requestBody interface{}, request []byte) (err error) {
	err = json.Unmarshal(request, &requestBody)
	return err
}

func (b *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	//extract isbn first
	isbn := r.URL.Query().Get("isbn")
	//call delete service
	rowsAffected, err := b.service.DeleteBook(isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Book %s deleted successfully (%d row affected)\n", isbn, rowsAffected)
}
