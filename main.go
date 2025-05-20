package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"orrett_backend/internal/db"

	"github.com/rs/cors"
)

type TotalBins struct {
	TotalBins int `json:"total_bins"`
}

func main() {
	fmt.Println("running")

	db.SetUp()

	mux := http.NewServeMux()

	mux.HandleFunc("/", getTotalBinsHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getTotalBinsHandler(w http.ResponseWriter, r *http.Request) {
	data := TotalBins{}
	data.TotalBins = 100 // Example value, replace with actual logic to get total bins
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
