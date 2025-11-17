package cart

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ordinaryteen/feez-go-api/internal/database"
	"github.com/ordinaryteen/feez-go-api/internal/middleware"
)

type AddToCartRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type CartItemResponse struct {
	ProductID    int64  `json:"product_id"`
	ProductName  string `json:"product_name"`
	Quantity     int    `json:"quantity"`
	PricePerItem int    `json:"price_per_item"`
}

func HandleGetCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	sqlQuery := `
		SELECT 
			c.product_id, 
			p.name, 
			c.quantity, 
			p.price
		FROM 
			public.cart_items c
		JOIN 
			public.products p ON c.product_id = p.id
		WHERE 
			c.user_id = $1
	`

	rows, err := database.DB.Query(context.Background(), sqlQuery, userID)
	if err != nil {
		http.Error(w, "Failed to fetch cart items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	cartItems := []CartItemResponse{}

	for rows.Next() {
		var item CartItemResponse
		if err := rows.Scan(
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.PricePerItem,
		); err != nil {
			http.Error(w, "Failed to process cart data", http.StatusInternalServerError)
			return
		}
		cartItems = append(cartItems, item)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading cart data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cartItems)
}

func HandleAddToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var req AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Quantity <= 0 {
		req.Quantity = 1
	}

	sqlQuery := `
		INSERT INTO public.cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`

	_, err := database.DB.Exec(
		context.Background(),
		sqlQuery,
		userID,
		req.ProductID,
		req.Quantity,
	)

	if err != nil {
		http.Error(w, "Failed to add item to cart: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Item added to cart"}`))
}
