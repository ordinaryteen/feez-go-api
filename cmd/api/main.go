package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	// import the  database
	"github.com/ordinaryteen/feez-go-api/internal/auth"
	_ "github.com/ordinaryteen/feez-go-api/internal/database"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, feez API! Server is running!"))
	})

	r.Post("/api/v1/signup", auth.HandleSignup)
	r.Post("/api/v1/login", auth.HandleLogin)

	port := ":8080"
	fmt.Println("Server listening on port", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
