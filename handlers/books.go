package handlers

import (
	"fmt"
	"go-labiblioteca-backend/database"
	"go-labiblioteca-backend/models"
	"log"
	"net/http"
)

type BookStore struct {
	bookList []string
	authors  []string
}

type BookHandler struct {
}

func (b *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//convert the string message to bytes
	var err error
	//sql.Open creates a pointer to Db (*Db)
	DbConn := database.ConnectDatabase()

	switch r.Method {
	case http.MethodGet:
		books := b.getBook(w)
		fmt.Println(books)

	}
	rows, err := DbConn.Query("SELECT * FROM books;")
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
}

func (b *BookHandler) getBook(w http.ResponseWriter) string {
	return "haha"
}

func (b *BookHandler) postBook(w http.ResponseWriter, r *http.Request) {
	//logic for post
}

func (b *BookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	//logic for update
}

func (b *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	//logic for delete
}
