package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

var films = map[string][]Film{
	"Films": {
		{Title: "The godfather", Director: "francis ford coppola"},
		{Title: "inception", Director: "christopher nolan"},
		{Title: "star wars", Director: "george lucas"},
	},
}

func h1(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, films)
}

func main() {
	fmt.Println(len(films))

	for key, asdf := range films {
		fmt.Println(key)
		fmt.Println(asdf)
	}

	http.HandleFunc("/", h1)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("failed to open http server %v", err)
	}
}
