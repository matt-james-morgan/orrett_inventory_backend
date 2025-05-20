package http

import (
	"encoding/json"
	"net/http"
	"orrett_backend/internal/service"

	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/totalBins", getTotalBinsHandler)
	mux.HandleFunc("/totalInventory", getTotalInventoryHandler)

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

func getTotalInventoryHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTotalInventory()
	if err != nil {
		http.Error(w, "Failed to fetch total inventory", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
