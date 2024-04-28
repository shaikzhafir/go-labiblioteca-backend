package main

import (
	"fmt"
	"labiblioteca/database"
	"labiblioteca/handlers"
	"labiblioteca/service"
	"labiblioteca/sqlcgen"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// adds the methods avail in bookModel to Env struct
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "main exit due to %s \n", err)
		os.Exit(1)
	}
}

func run() error {
	//init the db
	db, err := database.ConnectDatabase()
	if err != nil {
		return errors.Wrap(err, "database connection failed")
	}
	defer db.Close()
	//combo the dependencies here
	sqlcgen := sqlcgen.New(db)
	service := service.NewBookService(*sqlcgen)
	handler := handlers.NewBookHandler(service)

	mux := http.NewServeMux()
	// css and js files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// API routes
	mux.HandleFunc("GET /books", handler.GetBooks())
	mux.HandleFunc("POST /books", handler.InsertBook())
	mux.HandleFunc("GET /books/{id}", handler.GetBookByID())
	mux.HandleFunc("PUT /books/{id}", handler.UpdateBook())
	mux.HandleFunc("DELETE /books/{id}", handler.DeleteBook())

	// htmx routes
	mux.HandleFunc("GET /htmx/books", handler.GetBooksPage())

	// serve the html files
	mux.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Println("server is running on port 4000 yeet")

	err = http.ListenAndServe(":4000", mux)
	if err != nil {
		return errors.Wrap(err, "error when listening to server")
	}
	return nil
}
