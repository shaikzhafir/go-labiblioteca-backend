package handlers

import (
	"fmt"
	"go-labiblioteca-backend/data"
	"go-labiblioteca-backend/database"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type BookHandler struct {
	Message string
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

	bks := make([]data.Book, 0)
	for rows.Next() {
		bk := data.Book{}
		if err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Description, &bk.Author, &bk.Image); err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
		fmt.Println(bks)

	}
}

func (b *BookHandler) getBook(w http.ResponseWriter) string{
	return "haha"
}
