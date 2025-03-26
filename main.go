package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server on :8080")
	router := mux.NewRouter()

	// Handle root path "/"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the root route!")
	}).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed:", err)
	}
}