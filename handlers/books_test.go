//not incorporated yet, TODO

package handlers

import (
	"fmt"
	_ "net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
)

func TestBookHandler(t *testing.T) {
	bookHandler := &BookHandler{
		/* store: BookStore{
			bookList: []string{
				"poop",
				"haha",
			},
			authors: []string{
				"agatha",
				"John",
			},
		}, */
	}
	t.Run("Testing what the handler returns", func(t *testing.T) {
		//request, _ := http.NewRequest(http.MethodGet, "/books/John", nil)
		response := httptest.NewRecorder()

		//getBook accepts a response writer as its param
		got := bookHandler.getBook(response)
		want := "haha"

		if got != want {
			t.Errorf("expected %s but got %s", got, want)
			fmt.Println("lol")
		}

	})
}
