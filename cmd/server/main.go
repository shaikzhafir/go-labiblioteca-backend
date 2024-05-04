package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"labiblioteca/database"
	"labiblioteca/handlers"
	log "labiblioteca/logging"
	"labiblioteca/service"
	"labiblioteca/sqlcgen"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"github.com/dghubble/sessions"
	"github.com/pkg/errors"
	oauth2google "golang.org/x/oauth2/google"
)

var GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
var GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
var sessionCookieConfig *sessions.CookieConfig
var googleCookieConfig gologin.CookieConfig

var sessionStore sessions.Store[string]

const (
	sessionName     = "biblioteca-session"
	sessionUserKey  = "googleID"
	sessionUsername = "googleName"
	sessionEmail    = "googleEmail"
)

// sessionStore encodes and decodes session data stored in signed cookies
//var sessionStore = sessions.NewCookieStore[string](sessions.DebugCookieConfig, []byte(sessionSecret), nil)

const servingSchema = "http://"
const servingAddress = "localhost:4000"
const callbackPath = "/google/callback"

func init() {
	randKey, err := GenerateRandomKey(32)
	if err != nil {
		log.Fatal("Error generating random key")
	}
	if os.Getenv("PROD") == "true" {
		sessionCookieConfig = sessions.DefaultCookieConfig
	} else {
		sessionCookieConfig = sessions.DebugCookieConfig
	}
	sessionStore = sessions.NewCookieStore[string](sessionCookieConfig, randKey, nil)
}

// adds the methods avail in bookModel to Env struct
func main() {
	if len(GoogleClientID) == 0 || len(GoogleClientSecret) == 0 {
		log.Fatal("Set GOOGLE_CLIENT_* env vars")
	}
	if err := run(); err != nil {
		log.Fatal("Error running server: %v", err)
	}
}

// GenerateRandomKey generates a random key of specified length in bytes
func GenerateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	googleUser, err := google.UserFromContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session := sessionStore.New(sessionName)
	session.Set(sessionUserKey, googleUser.Id)
	session.Set(sessionUsername, googleUser.Name)
	session.Set(sessionEmail, googleUser.Email)
	if err := session.Save(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func renderTemplate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, sessionName)
		if err != nil {
			tmpl, err := template.ParseFiles("public/index.html")
			if err != nil {
				log.Error("Error rendering template: %v", err)
				http.Error(w, "Error loading template", http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				log.Error("Error rendering template: %v", err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}

		tmpl, err := template.ParseFiles("templates/layout.tmpl", "templates/profile.tmpl")
		if err != nil {
			log.Error("Error rendering template: %v", err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		log.Info("session: %v", session)
		err = tmpl.ExecuteTemplate(w, "layout.tmpl", map[string]string{
			"Name":  session.Get(sessionUsername),
			"Email": session.Get(sessionEmail),
		})
		if err != nil {
			log.Error("Error rendering template: %v", err)
			http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		}
	})
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

	conf := &oauth2.Config{
		ClientID:     GoogleClientID,
		ClientSecret: GoogleClientSecret,
		RedirectURL:  servingSchema + servingAddress + callbackPath,
		Scopes:       []string{"profile", "email"},
		Endpoint:     oauth2google.Endpoint,
	}
	if os.Getenv("PROD") == "true" {
		googleCookieConfig = gologin.DefaultCookieConfig
	} else {
		googleCookieConfig = gologin.DebugOnlyCookieConfig
	}

	mux := http.NewServeMux()
	// css and js files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// API routes
	mux.HandleFunc("GET /books", handler.GetBooks())
	mux.HandleFunc("POST /books", handler.InsertBook())
	mux.HandleFunc("GET /books/{id}", handler.GetBookByID())
	mux.HandleFunc("PUT /books/{id}", handler.UpdateBook())
	mux.HandleFunc("DELETE /books/{id}", handler.DeleteBook())
	mux.Handle("/login/", google.StateHandler(googleCookieConfig, google.LoginHandler(conf, nil)))
	mux.Handle(callbackPath, google.StateHandler(googleCookieConfig, google.CallbackHandler(conf, http.HandlerFunc(googleCallbackHandler), nil)))
	mux.Handle("POST /logout", http.HandlerFunc(logoutHandler))

	// htmx routes
	mux.HandleFunc("GET /htmx/books", handler.GetBooksPage())

	// home page routing
	mux.Handle("/", renderTemplate())

	fmt.Println("server is running on port 4000 yeet")
	address := ":4000"
	if os.Getenv("PROD") == "true" {
		address = os.Getenv("PROD_ADDRESS")
	}
	err = http.ListenAndServe(address, mux)
	if err != nil {
		return errors.Wrap(err, "error when listening to server")
	}
	return nil
}

// logoutHandler destroys the session on POSTs and redirects to home.
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	sessionStore.Destroy(w, sessionName)
	http.Redirect(w, req, "/", http.StatusFound)
}
