package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Server", "Go")

	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid snippet ID", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Display a form for creating a new snippet..."))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetCreatePost(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)

	_, err := w.Write([]byte("Save a new snippet..."))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
