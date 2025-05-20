package service

import (
	"orrett_backend/internal/db"
	"orrett_backend/internal/models"
)

func GetTotalBins() ([]models.TotalBins, error) {
	return db.FetchTotalBins()
}
