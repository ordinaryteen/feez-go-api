package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	// import the  database
	"github.com/ordinaryteen/feez-go-api/internal/auth"
	"github.com/ordinaryteen/feez-go-api/internal/cart"
	_ "github.com/ordinaryteen/feez-go-api/internal/database"
	"github.com/ordinaryteen/feez-go-api/internal/middleware"
	"github.com/ordinaryteen/feez-go-api/internal/product"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, feez API! Server is running!"))
	})

	// --- API v1 (Public Routes) ---
	r.Post("/api/v1/signup", auth.HandleSignup)
	r.Post("/api/v1/login", auth.HandleLogin)
	r.Get("/api/v1/products", product.HandleGetProducts)

	// --- API v1 (Private Routes) ---
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/api/v1/cart", cart.HandleAddToCart)
	})

	port := ":8080"
	fmt.Println("Server listening on port", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
