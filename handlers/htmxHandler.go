package handlers

import (
	"fmt"
	"html/template"
	log "labiblioteca/logging"
	"net/http"
)

func (b *BookHandler) GetBooksPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("htmx get books called")
		bks, err := b.service.GetBooks(r.Context())
		if err != nil {
			fmt.Println("haha")
		}
		// return some html content here
		tmpl, err := template.ParseFiles("templates/bookEntry.tmpl")
		if err != nil {
			log.Error(err.Error())
			w.Write([]byte("error parsing template"))
			return
		}
		for _, bk := range bks {
			tmpl.Execute(w, bk)
		}
	}
}
