package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	msg := fmt.Sprintf("Snippet view: %d", id)

	_, writeErr := w.Write([]byte(msg))
	if writeErr != nil {
		log.Printf("failed to write response: %v", writeErr)
	}
}

func snippetCreateView(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreateView)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
