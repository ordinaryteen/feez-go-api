package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/ordinaryteen/feez-go-api/internal/database"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var userID string
	sqlQuery := `
		SELECT id FROM private.auth_users
		WHERE email = $1 AND hashed_password = crypt($2, hashed_password)
	`

	err := database.DB.QueryRow(
		context.Background(),
		sqlQuery,
		req.Email,
		req.Password,
	).Scan(&userID)

	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Failed to login: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokenString, err := GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
