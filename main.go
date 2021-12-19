package main

import (
	_ "fmt"
	"go-labiblioteca-backend/database"
	"go-labiblioteca-backend/models"
	"go-labiblioteca-backend/handlers"
	"log"
	"net/http"
	_ "os"

	"github.com/joho/godotenv"
)

//adds the methods avail in bookModel to Env struct
type Env struct {
	books models.BookModel
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading env")
	}
	//init the db
	db,err := database.ConnectDatabase()
	if err != nil {
		log.Fatalln("Error in connecting db")
	}

	//allows the db instance to be used for models
	env := Env{
		books: models.BookModel{DB: db},
	}
	//initialises the db if doesnt exist
	env.books.CreateBookTable()

	handler := handlers.BookHandler{DB : db}
	//test ur env variables lol
	/* 	fmt.Println("Shell:", os.Getenv("TEST"))
	   	test := os.Getenv("TEST")
	   	fmt.Println(test)
	*/
	
	http.HandleFunc("/books", handler.ServeHTTP)
	err = http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatalln("Error in listening to port")
	}

}


