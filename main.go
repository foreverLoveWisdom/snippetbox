package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func snippetView(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Display a specific snippet..."))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func snippetCreateView(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func exampleHandler(_ http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	itemID := r.PathValue("itemID")
	log.Printf("category: %s, itemID: %s", category, itemID)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreateView)
	mux.HandleFunc("/products/{category}/items/{itemID}/", exampleHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
