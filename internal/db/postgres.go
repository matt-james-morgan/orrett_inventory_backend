package db

import (
	"database/sql"
	"orrett_backend/internal/models"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func FetchTotalBins() ([]models.TotalBins, error) {
	rows, err := db.Query("SELECT id, total FROM bins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bins []models.TotalBins
	for rows.Next() {
		var bin models.TotalBins
		if err := rows.Scan(&bin.ID, &bin.Total); err != nil {
			return nil, err
		}
		bins = append(bins, bin)
	}
	return bins, nil
}
