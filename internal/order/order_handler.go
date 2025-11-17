package order

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ordinaryteen/feez-go-api/internal/database"
	"github.com/ordinaryteen/feez-go-api/internal/middleware"
)

type CheckoutResponse struct {
	OrderID int64 `json:"order_id"`
}

func HandleCheckout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var newOrderID int64

	sqlQuery := `SELECT public.checkout($1)`

	err := database.DB.QueryRow(
		context.Background(),
		sqlQuery,
		userID,
	).Scan(&newOrderID)

	if err != nil {
		http.Error(w, "Checkout failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CheckoutResponse{OrderID: newOrderID})
}
