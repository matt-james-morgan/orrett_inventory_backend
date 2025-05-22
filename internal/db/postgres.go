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
func FetchBins() ([]models.Bin, error) {
	rows, err := db.Query("SELECT bin_name, id FROM bins")
	if err != nil {
		return []models.Bin{}, err
	}
	fmt.Println(rows, "rows")
	defer rows.Close()

	var bins []models.Bin

	for rows.Next() {
		var bin models.Bin
		if err := rows.Scan(&bin.Name, &bin.ID); err != nil {
			return []models.Bin{}, err
		}
		bins = append(bins, bin)
	}

	return bins, nil
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

func CreateBin(binName string) (models.Bin, error) {
	const query = `INSERT INTO bins (bin_name) VALUES ($1) RETURNING id, bin_name`
	var bin models.Bin
	err := db.QueryRow(query, binName).Scan(&bin.ID, &bin.Name)
	if err != nil {
		return models.Bin{}, fmt.Errorf("failed to insert bin: %w", err)
	}
	return bin, nil
}
