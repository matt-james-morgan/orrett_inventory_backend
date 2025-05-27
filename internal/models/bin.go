package models

type TotalBins struct {
	Total int `json:"totalBins"`
}

type TotalItems struct {
	Total int `json:"totalItems"`
}

type Bin struct {
	Name        string `json:"binName"`
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type Item struct {
	ItemName    string `json:"itemName"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	BinName     string `json:"binName"`
}
