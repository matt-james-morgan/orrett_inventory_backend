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

func CreateBin(bin_name, description string) (models.Bin, error) {
	return db.CreateBin(bin_name, description)
}

func CreateItem(item_name, description, bin_id string) (models.Item, error) {
	return db.CreateItem(item_name, description, bin_id)
}
