package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orrett_backend/internal/service"
)

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTotalItems()
	if err != nil {
		http.Error(w, "Failed to fetch total items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	type CreateItemRequest struct {
		ItemName string `json:"itemName"`
		BinId    int    `json:"binId"`
	}

	var req CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.BinId == 0 {
		http.Error(w, "bin_id is required", http.StatusBadRequest)
		return
	}

	if req.ItemName == "" {
		http.Error(w, "item_name is required", http.StatusBadRequest)
		return
	}

	data, err := service.CreateItem(req.ItemName, req.BinId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	type DeleteItemRequest struct {
		ItemId int `json:"itemId"`
	}

	var req DeleteItemRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.ItemId == 0 {
		http.Error(w, "Invalid Item ID", http.StatusBadRequest)
		return
	}

	ok, err := service.DeleteItem(req.ItemId)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"success": ok}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
