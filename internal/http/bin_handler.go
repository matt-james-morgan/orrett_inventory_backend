package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orrett_backend/internal/service"
)

func GetTotalBinsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTotalBins()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch total bins", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func CreateBin(w http.ResponseWriter, r *http.Request) {
	type CreateBinRequest struct {
		BinName     string `json:"binName"`
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

func DeleteBin(w http.ResponseWriter, r *http.Request) {
	type DeleteBinRequest struct {
		BinId int `json:"BinId"`
	}

	var req DeleteBinRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.BinId == 0 {
		http.Error(w, "Invalid Bin ID", http.StatusBadRequest)
		return
	}

	ok, err := service.DeleteBin(req.BinId)
	if err != nil {
		http.Error(w, "Failed to delete Bin", http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"success": ok}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
