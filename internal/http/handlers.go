package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orrett_backend/internal/service"

	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/bins", getTotalBinsHandler)
	mux.HandleFunc("/totalInventory", getBinsHandler)
	mux.HandleFunc("/createBin", createBin)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	return handler
}

func getTotalBinsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTotalBins()
	if err != nil {
		http.Error(w, "Failed to fetch total bins", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func getBinsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTotalBins()
	if err != nil {
		http.Error(w, "Failed to fetch total inventory", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func createBin(w http.ResponseWriter, r *http.Request) {
	type CreateBinRequest struct {
		BinName     string `json:"bin_name"`
		Description string `json:"description"`
	}

	var req CreateBinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.BinName == "" {
		http.Error(w, "bin_name is required", http.StatusBadRequest)
		return
	}

	data, err := service.CreateBin(req.BinName, req.Description)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create bin", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}
