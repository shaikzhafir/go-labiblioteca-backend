package main

import (
	"encoding/json"
	_ "fmt"
	"go-labiblioteca-backend/database"
	"go-labiblioteca-backend/models"
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
	db := database.ConnectDatabase()

	//allows the db instance to be used for models
	env := Env{
		books: models.BookModel{DB: db},
	}
	//test ur env variables lol
	/* 	fmt.Println("Shell:", os.Getenv("TEST"))
	   	test := os.Getenv("TEST")
	   	fmt.Println(test)
	*/
	//initialises the db if doesnt exist
	env.books.CreateBookTable()
	http.HandleFunc("/books", env.getBooks)
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalln("Error in listening to port")
	}

}

//refactor to put handlers in proper place
func (env *Env) getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := env.books.DB.Query("SELECT * FROM books;")
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
	}

	//converts book array over to JSON and addes to reponsewriter
	if err := json.NewEncoder(w).Encode(bks); err != nil {
		log.Println(err)
	}
}
