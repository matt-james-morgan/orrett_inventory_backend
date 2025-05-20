package db

import (
	"database/sql"
	"fmt"
	"orrett_backend/internal/models"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}
func FetchTotalBins() (models.TotalBins, error) {
	rows, err := db.Query("SELECT COUNT(id) as total FROM bins")
	if err != nil {
		return models.TotalBins{}, err
	}
	defer rows.Close()

	var binTotal models.TotalBins

	if rows.Next() {
		if err := rows.Scan(&binTotal.Total); err != nil {
			return models.TotalBins{}, err
		}
	} else {
		return models.TotalBins{}, fmt.Errorf("no rows returned")
	}

	return binTotal, nil
}

func FetchTotalInventory() (models.TotalInventory, error) {
	rows, err := db.Query("SELECT COUNT(item_name) as total FROM inventory")
	if err != nil {
		return models.TotalInventory{}, err
	}
	defer rows.Close()

	var inventoryTotal models.TotalInventory

	if rows.Next() {
		if err := rows.Scan(&inventoryTotal.Total); err != nil {
			return models.TotalInventory{}, err
		}
	} else {
		return models.TotalInventory{}, fmt.Errorf("no rows returned")
	}

	return inventoryTotal, nil
}
