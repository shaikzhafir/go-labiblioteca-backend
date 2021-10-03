package main


import (
	"net/http"
	"fmt"
)



func BookServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Stoner") 
}