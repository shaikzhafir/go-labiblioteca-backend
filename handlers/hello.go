package handlers

import (
	"fmt"
	"go-labiblioteca-backend/data"
	"go-labiblioteca-backend/database"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type HelloHandler struct {
	Message string
}

// Hello this is just to return some message
func Hello() *HelloHandler {
	return &HelloHandler{Message: "POOP"}
}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//convert the string message to bytes
	var err error
	//sql.Open creates a pointer to Db (*Db)
	DbConn := database.ConnectDatabase()
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
