package service

import (
	"orrett_backend/internal/db"
	"orrett_backend/internal/models"
)

func GetTotalBins() ([]models.Bin, error) {
	return db.FetchBins()
}

func GetTotalInventory() (models.TotalInventory, error) {
	return db.FetchTotalInventory()
}

func CreateBin(bin_name string) (models.Bin, error) {
	return db.CreateBin(bin_name)
}
