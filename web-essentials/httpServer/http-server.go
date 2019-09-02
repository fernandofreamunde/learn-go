package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

		// ok I didnt get the file server thing yet, more work for later will carry on to the next lesson
		fmt.Fprintf(w, "Welcome to my site! <img src=\"static/img.jpg\">")
	})

	fs := http.FileServer(http.Dir("/static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":81", nil)
}
