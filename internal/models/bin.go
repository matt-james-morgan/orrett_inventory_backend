package models

type TotalBins struct {
	Total int `json:"totalBins"`
}

type TotalItems struct {
	Total int `json:"totalItems"`
}

type Bin struct {
	Name        string `json:"name"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
}

type Item struct {
	Name        string `json:"name"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	BinId       string `json:"binId"`
}
