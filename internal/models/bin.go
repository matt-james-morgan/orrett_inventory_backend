package models

type TotalBins struct {
	Total int `json:"total_bins"`
}

type TotalInventory struct {
	Total int `json:"total_inventory"`
}

type Bin struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
