package service

import (
	"orrett_backend/internal/db"
	"orrett_backend/internal/models"
)

func GetTotalBins() ([]models.Bin, error) {
	return db.FetchBins()
}

func GetTotalItems() (models.TotalItems, error) {
	return db.FetchTotalItems()
}

func CreateBin(bin_name, description string) (models.Bin, error) {
	return db.CreateBin(bin_name, description)
}

func CreateItem(item_name string, bin_id int) (models.Item, error) {
	return db.CreateItem(item_name, bin_id)
}

func DeleteItem(item_id int) (bool, error) {
	return db.DeleteItem(item_id)
}

func DeleteBin(bin_id int) (bool, error) {
	return db.DeleteBin(bin_id)
}

func SignIn(user_name string) (string, error) {
	return db.SignIn(user_name)
}
