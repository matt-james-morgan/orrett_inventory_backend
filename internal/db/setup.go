package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "password"
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "orrett_inventory"
	dbTable    = "inventory"
)

func SetUp() *sql.DB {
	// Connect to default "postgres" DB to create new DB
	defaultDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword)

	defaultDB, err := sql.Open("postgres", defaultDSN)
	if err != nil {
		log.Fatalf("Failed to connect to default DB: %v", err)
	}
	defer defaultDB.Close()

	// Connect to the newly created DB
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to new DB: %v", err)
	}

	return db
}
