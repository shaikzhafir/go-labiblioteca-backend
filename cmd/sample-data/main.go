//logic to add sample books to db to initliase

package main

import (
	"go-labiblioteca-backend/database"
	"go-labiblioteca-backend/domain"
	"go-labiblioteca-backend/repository"
	"log"
)

var sampleBooks = []domain.Book{
	{
		Isbn:        "random isbn",
		Title:       "asdasd",
		Description: "asdasd",
		Author:      "asd",
		ImageURL:    "asdasd",
	},
	{
		Isbn:        "asd",
		Title:       "asd",
		Description: "weqwe",
		Author:      "qweqwe",
		ImageURL:    "qweqwe",
	},
	{
		Isbn:        "qweqweqwe",
		Title:       "qweqwe",
		Description: "qweqwe",
		Author:      "qwe",
		ImageURL:    "qwe",
	},
	{
		Isbn:        "zxczxc",
		Title:       "zxczxc",
		Description: "zxczxc",
		Author:      "zxczxc",
		ImageURL:    "zxczxc",
	},
	{
		Isbn:        "cvbcvb",
		Title:       "cvbcvb",
		Description: "cvbcvb",
		Author:      "cvbcvb",
		ImageURL:    "cvbcvb",
	},
}

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("error connecting to db %s", err)
	}
	repo := repository.NewBookRepository(db)
	repo.CreateBookTable()
	err = repo.InsertManyBooks(&sampleBooks)
	if err != nil {
		log.Fatalf("books not seeded successfully, error occured %s", err)
	}
	log.Printf("database successfully seeded\n")
}
