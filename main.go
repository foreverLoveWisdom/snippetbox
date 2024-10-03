package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Server", "Go HTTP Server")

	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreateView(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func snippetCreatePost(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)

	_, err := w.Write([]byte("Save new snippet..."))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreateView)

	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
