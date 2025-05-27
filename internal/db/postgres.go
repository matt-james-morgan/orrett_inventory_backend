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
	rows, err := db.Query("SELECT bin_name, description, id FROM bins")
	if err != nil {
		return []models.Bin{}, err
	}
	defer rows.Close()

	var bins []models.Bin

	for rows.Next() {
		var bin models.Bin
		if err := rows.Scan(&bin.Name, &bin.Description, &bin.ID); err != nil {
			return []models.Bin{}, err
		}
		bins = append(bins, bin)
	}

	return bins, nil
}

func FetchTotalItems() (models.TotalItems, error) {
	rows, err := db.Query("SELECT COUNT(item_name) as total FROM items")
	if err != nil {
		return models.TotalItems{}, err
	}
	defer rows.Close()

	var itemsTotal models.TotalItems

	if rows.Next() {
		if err := rows.Scan(&itemsTotal.Total); err != nil {
			return models.TotalItems{}, err
		}
	} else {
		return models.TotalItems{}, fmt.Errorf("no rows returned")
	}

	return itemsTotal, nil
}

func CreateBin(binName, description string) (models.Bin, error) {
	const query = `INSERT INTO bins (bin_name, description) VALUES ($1, $2) RETURNING id, bin_name, description`
	var bin models.Bin
	err := db.QueryRow(query, binName, description).Scan(&bin.ID, &bin.Name, &bin.Description)
	if err != nil {
		return models.Bin{}, fmt.Errorf("failed to insert bin: %w", err)
	}
	return bin, nil
}

func CreateItem(item_name, description, bin_name string) (models.Item, error) {
	const query = `INSERT INTO items (item_name, bin_id, description) VALUES ($1, $2, $3) RETURNING id, item_name, bin_id, description`
	var item models.Item
	err := db.QueryRow(query, item_name, bin_name, description).Scan(&item.ID, &item.ItemName, &item.BinName, &item.Description)
	if err != nil {
		return models.Item{}, fmt.Errorf("failed to insert bin: %w", err)
	}
	return item, nil
}
