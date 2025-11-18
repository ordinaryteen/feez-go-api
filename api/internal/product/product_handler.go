package product

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ordinaryteen/feez-go-api/internal/database"
)

// Product defines the structure for product data
type Product struct {
	ID               int64  `json:"id"`
	CategoryID       int64  `json:"category_id"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
	StockTersisa     int    `json:"stock_tersisa"`
	CreatedByAdminID string `json:"created_by_admin_id"` // UUID as string
}

// HandleGetProducts fetches all products from the database
func HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	sqlQuery := `
		SELECT id, category_id, name, price, stock_tersisa, created_by_admin_id 
		FROM public.products
		ORDER BY name ASC
	`
	rows, err := database.DB.Query(context.Background(), sqlQuery)
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ID,
			&p.CategoryID,
			&p.Name,
			&p.Price,
			&p.StockTersisa,
			&p.CreatedByAdminID,
		); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			http.Error(w, "Failed to process product data", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v\n", err)
		http.Error(w, "Error reading product data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
