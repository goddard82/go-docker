package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/searcher", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hi")
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
