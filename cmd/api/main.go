package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	// import the  database
	_ "github.com/ordinaryteen/feez-go-api/internal/database"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, feez API! Server is running!"))
	})

	port := ":8080"
	fmt.Println("Server listening on port", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
