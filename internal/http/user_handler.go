package http

import (
	"encoding/json"
	"net/http"
	"orrett_backend/internal/service"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check username + password
	passwordHash, err := service.SignIn(req.Username)

	if err != nil {
		http.Error(w, "Failed to sign in", http.StatusInternalServerError)
		return
	}
	ok := CheckPasswordHash(req.Password, passwordHash)
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// On success, return JSON (for now, just a simple success message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
