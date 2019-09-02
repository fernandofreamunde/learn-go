package main

// requires to run: go get -u github.com/gorilla/mux
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/book/{title}/page/{page}", func (w http.ResponseWriter, r *http.Request) {

		vars  := mux.Vars(r) //r for the request... these tutorials using bad var names are going to kill me xD
		title := vars["title"]
		page  := vars["page"]

		fmt.Fprintf(w, "You have requested the book: %s on the page %s\n", title, page)
	}).Methods("GET")

	http.ListenAndServe(":81", router)
}