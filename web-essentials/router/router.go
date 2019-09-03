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

// ### Features of the gorilla/mux Router
// # Methods
// Restrict the request handler to specific HTTP methods.

// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// # Hostnames & Subdomains
// Restrict the request handler to specific hostnames or subdomains.

// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

// # chemes
// Restrict the request handler to http/https.

// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

// # Path Prefixes & Subrouters
// Restrict the request handler to specific path prefixes.

// bookrouter := r.PathPrefix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBooks)
// bookrouter.HandleFunc("/{title}", GetBook)
