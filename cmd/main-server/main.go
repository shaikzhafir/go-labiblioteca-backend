package main

import (
	"fmt"
	_ "fmt"
	"go-labiblioteca-backend/database"
	"go-labiblioteca-backend/handlers"
	"go-labiblioteca-backend/repository"
	"go-labiblioteca-backend/service"
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

// server is used to implement helloworld.GreeterServer.
/* type server struct {
	pb.UnimplementedBookRepoServer
}
*/
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
	// this is some useless grpc stuff
	/* var (
		port = flag.Int("port", 50051, "The server port")
	)
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookRepoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} */

	http.HandleFunc("/books", handler.ServeHTTP)
	err = http.ListenAndServe("0.0.0.0:4000", nil)
	if err != nil {
		return errors.Wrap(err, "error when listening to server")
	}

	defer db.Close()
	return nil
}
