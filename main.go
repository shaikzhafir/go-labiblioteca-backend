package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"go-labiblioteca-backend/handlers"
)



func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading env")
	}
	fmt.Println("Shell:", os.Getenv("TEST"))
	test := os.Getenv("TEST")
	fmt.Println(test)
	http.Handle("/",&handlers.HelloHandler{
		Message: "yohoho",
	})

	err = http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatalln("Error in listening to port")
	}
}
