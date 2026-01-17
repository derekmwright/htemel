package main

import (
	"net/http"

	"github.com/derekmwright/htemel/examples/webapp/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetRoot)
	mux.HandleFunc("/about", handlers.GetAbout)
	mux.HandleFunc("/contact", handlers.GetContact)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
