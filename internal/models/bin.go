package models

type TotalBins struct {
	Total int `json:"total_bins"`
}

type TotalInventory struct {
	Total int `json:"total_inventory"`
}

type Bin struct {
	Name        string `json:"bin_name"`
	ID          int    `json:"id"`
	Description string `json:"description"`
}
