package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"labiblioteca/domain"
	log "labiblioteca/logging"
	"labiblioteca/service"
	"labiblioteca/sqlcgen"
	"net/http"
	"strconv"
	_ "strings"
)

type BookHandler struct {
	service *service.BookService
}

// NewBookHandler for it to be called by outer domain
func NewBookHandler(service *service.BookService) BookHandler {
	return BookHandler{service: service}
}

func (b *BookHandler) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("GetBooks called")
		bks, err := b.service.GetBooks(r.Context())
		if err != nil {
			fmt.Println("haha")
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(bks); err != nil {
			log.Error(err.Error())
		}
	}
}

func (b *BookHandler) GetBookByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//extract id first
		idString := r.PathValue("id")
		// convert to int
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bks, err := b.service.GetBookByID(r.Context(), id)
		if err != nil {
			fmt.Println("haha")
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(bks); err != nil {
			log.Error(err.Error())
		}
	}
}

func (b *BookHandler) InsertBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var insertParams sqlcgen.InsertBookParams
		err := json.NewDecoder(r.Body).Decode(&insertParams)
		//handle error for decoding
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("insert params %v", insertParams)
		err = b.service.AddBook(r.Context(), &insertParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Book %s created successfully", insertParams.Title)
	}
}

func (b *BookHandler) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//extract id first
		idString := r.PathValue("id")
		// convert to int
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book, err := unmarshalBook(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = b.service.UpdateBook(r.Context(), &sqlcgen.UpdateBookParams{
			Isbn:        book.Isbn,
			Title:       book.Title,
			Description: book.Description,
			ImageUrl:    book.ImageURL,
			ID:          id,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func unmarshalBook(r *http.Request) (book domain.Book, err error) {
	data, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(data, &book)
	return book, err
}

func (b *BookHandler) DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//extract id first
		idString := r.PathValue("id")
		// convert to int
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//call delete service
		err = b.service.DeleteBook(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
