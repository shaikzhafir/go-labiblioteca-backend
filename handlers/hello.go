package handlers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-labiblioteca-backend/data"
	"log"
	"net/http"
)


type HelloHandler struct {
	Message string
}

// Hello this is just to return some message
func Hello() *HelloHandler {
	return &HelloHandler{Message: "POOP"}
}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	//convert the string message to bytes
	var err error
	//sql.Open creates a pointer to Db (*Db)
	DbConn, err := sql.Open("postgres", "postgres://postgres:poop@localhost/bookdb?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	DbConn.SetMaxOpenConns(3)
	rows,err := DbConn.Query("SELECT * FROM books;")
	if err != nil {
		log.Fatal(err)
	}

	bks := make([]data.Book,0)
	for rows.Next() {
		bk := data.Book{}
		if err = rows.Scan(&bk.Isbn,&bk.Title,&bk.Description,&bk.Author,&bk.Image);err != nil {
			log.Fatal(err)
		}
		bks = append(bks,bk)
		fmt.Println(bks)

	}
	}
