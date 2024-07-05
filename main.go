package main

import (
	"html/template"
	"log"
	"net/http"
)

func h1(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", h1)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("failed to open http server %v", err)
	}
}
