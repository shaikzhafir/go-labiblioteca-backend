package handlers

import (
	"fmt"
	"database/sql"
	"go-labiblioteca-backend/models"
	"log"
	"net/http"
	"encoding/json"
)

type BookStore struct {
	bookList []string
	authors  []string
}

//when initialising the handler struct, include the db connection 
// that came in main.go
type BookHandler struct {
	DB *sql.DB 
}

func (b *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		b.getBook(w)
	case http.MethodPost:
		b.postBook(w,r)
	case http.MethodPut:
		b.updateBook(w,r)
	case http.MethodDelete:
		b.deleteBook(w,r)
	}
		
}

func (b *BookHandler) getBook(w http.ResponseWriter) {
	rows, err := b.DB.Query("SELECT * FROM books;")
	if err != nil {
		log.Fatal(err)
	}

	bks := make([]models.Book, 0)
	for rows.Next() {
		bk := models.Book{}
		if err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Description, &bk.Author, &bk.ImageURL); err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
		fmt.Println(bks)
	}
	
	if err := json.NewEncoder(w).Encode(bks); err != nil {
		log.Println(err)
	}
}

func (b *BookHandler) postBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result,err := b.DB.Exec("INSERT INTO books(isbn, title, description, author, imageURL) VALUES($1, $2, $3, $4, $5);", &book.Isbn, &book.Title, &book.Description, &book.Author, &book.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected,err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", book.Title, rowsAffected)
	}


func (b *BookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	//logic for update
}

func (b *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	//logic for delete

	//extract isbn first 
	isbn := r.URL.Query().Get("isbn")

	//execute delete 
	result,err := b.DB.Exec("DELETE FROM books WHERE isbn = $1",isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected,err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Book %s deleted successfully (%d row affected)\n", isbn, rowsAffected)
}
