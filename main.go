package main

import (
	"log"
	"net/http"
	"orrett_backend/internal/db"
	handlers "orrett_backend/internal/http"
)

type TotalBins struct {
	TotalBins int `json:"total_bins"`
}

func main() {

	dbConn := db.SetUp() // your SetUp returns *sql.DB
	defer dbConn.Close()

	db.SetDB(dbConn)

	log.Println("Query executed")

	handler := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handler))
}
