package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGETBooks(t *testing.T){
	t.Run("Returns author names", func (t *testing.T){
		request, _ := http.NewRequest(http.MethodGet, "/books/John", nil)
		response := httptest.NewRecorder() 
	
		BookServer(response,request)
	
		got := response.Body.String()
		want := "Stoner"
	
		if got != want {
			t.Errorf("got %s, want %s",got,want)
		}
	})
}
