package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"orrett_backend/internal/models"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func FetchBins() ([]models.Bin, error) {
	rows, err := db.Query(
		`SELECT 
    		bins.id AS bin_id, 
    		bins.bin_name, 
    		bins.description, 
    		items.id AS item_id, 
    		items.item_name,
		FROM bins
		LEFT JOIN items ON bins.id = items.bin_id
		ORDER BY bins.id
		`)
	if err != nil {
		slog.Warn("error forming query")
		return nil, err
	}
	defer rows.Close()

	binsMap := make(map[int]*models.Bin)

	for rows.Next() {
		var binID int
		var binName string
		var description string
		var itemID sql.NullInt64
		var itemName sql.NullString

		if err := rows.Scan(&binID, &binName, &description, &itemID, &itemName); err != nil {
			slog.Warn("error scanning")
			fmt.Println("error scanning")
			return []models.Bin{}, err
		}

		bin, exists := binsMap[binID]
		if !exists {
			bin = &models.Bin{
				ID:          binID,
				Name:        binName,
				Description: description,
				Items:       []models.Item{},
			}
			binsMap[binID] = bin
		}

		if itemID.Valid {
			bin.Items = append(bin.Items, models.Item{
				ID:       int(itemID.Int64),
				ItemName: itemName.String,
			})
		}
	}

	var bins []models.Bin

	for _, bin := range binsMap {
		bins = append(bins, *bin)
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

func CreateItem(item_name, bin_name string) (models.Item, error) {
	const query = `INSERT INTO items (item_name, bin_id) VALUES ($1, $2, $3) RETURNING id, item_name, bin_id`
	var item models.Item
	err := db.QueryRow(query, item_name, bin_name).Scan(&item.ID, &item.ItemName)
	if err != nil {
		return models.Item{}, fmt.Errorf("failed to insert bin: %w", err)
	}
	return item, nil
}
