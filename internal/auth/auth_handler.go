package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ordinaryteen/feez-go-api/internal/database"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sqlQuery := `SELECT public.signup($1, $2, $3)`

	_, err := database.DB.Exec(
		context.Background(),
		sqlQuery,
		req.Email,
		req.Password,
		req.Username,
	)

	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Signup successful"}`))
}
