package main

import (
	"fmt"
	_ "fmt"
	"github.com/pkg/errors"
	"go-labiblioteca-backend/database"
	"go-labiblioteca-backend/handlers"
	"go-labiblioteca-backend/repository"
	"go-labiblioteca-backend/service"
	"net/http"
	"os"
)

//adds the methods avail in bookModel to Env struct
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
	//combo the dependencies here
	repo := repository.NewBookRepository(db)
	repo.CreateBookTable()
	service := service.NewBookService(repo)
	handler := handlers.NewBookHandler(service)
	//test ur env variables lol
	/* 	fmt.Println("Shell:", os.Getenv("TEST"))
	   	test := os.Getenv("TEST")
	   	fmt.Println(test)
	*/

	http.HandleFunc("/books", handler.ServeHTTP)
	err = http.ListenAndServe("0.0.0.0:4000", nil)
	if err != nil {
		return errors.Wrap(err, "error when listening to server")
	}

	defer db.Close()
	return nil
}
