package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "task 1", Done: true},
				{Title: "task 2", Done: true},
				{Title: "task 3", Done: false},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":81", nil)
}
