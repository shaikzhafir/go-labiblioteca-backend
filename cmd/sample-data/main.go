//logic to add sample books to db to initliase

package main

import (
	"context"
	"labiblioteca/database"
	"labiblioteca/sqlcgen"
	"log"
)

var sampleBooks = []sqlcgen.InsertBookParams{
	{
		Isbn:        "random isbn",
		Title:       "asdasd",
		Description: "asdasd",
		Author:      "asd",
		ImageUrl:    "asd",
	},
	{
		Isbn:        "random isbn",
		Title:       "asdasd",
		Description: "asdasd",
		Author:      "asd",
		ImageUrl:    "asd",
	},
	{
		Isbn:        "random isbn",
		Title:       "asdasd",
		Description: "asdasd",
		Author:      "asd",
		ImageUrl:    "asd",
	},
	{
		Isbn:        "random isbn",
		Title:       "asdasd",
		Description: "asdasd",
		Author:      "asd",
		ImageUrl:    "asd",
	},
	{
		Isbn:        "random isbn",
		Title:       "asdasd",
		Description: "asdasd",
		Author:      "asd",
		ImageUrl:    "asd",
	},
}

func main() {
	ctx := context.Background()
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("error connecting to db %s", err)
	}

	sqlc := sqlcgen.New(db)
	// create tables if not exists
	result, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn TEXT, title TEXT, description TEXT, author TEXT, image_url TEXT);")
	if err != nil {
		log.Fatalf("error creating table %s", err)
	}
	log.Println(result)
	for _, book := range sampleBooks {

		err = sqlc.InsertBook(ctx, book)
		if err != nil {
			log.Fatalf("error inserting book %s", err)
		}
	}
}
