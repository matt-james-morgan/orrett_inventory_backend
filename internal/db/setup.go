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
	dbName     = "orrett_iventory"
	dbTable    = "inventory"
)

func SetUp() {
	// Connect to default "postgres" DB to create new DB
	defaultDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword)

	defaultDB, err := sql.Open("postgres", defaultDSN)
	if err != nil {
		log.Fatalf("Failed to connect to default DB: %v", err)
	}
	defer defaultDB.Close()

	// Try to create the new database
	_, err = defaultDB.Exec("CREATE DATABASE " + dbName)
	if err != nil && !isDuplicateDBError(err) {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Now connect to the newly created DB
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to new DB: %v", err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS inventory (
		item_name TEXT,
		bin_id INTEGER
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Insert sample data
	_, err = db.Exec(`INSERT INTO inventory (item_name, bin_id) VALUES 
		('Screwdriver', 101), 
		('Hammer', 102), 
		('Wrench', 101)`)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	log.Println("Database, table, and sample data set up successfully.")
}

func isDuplicateDBError(err error) bool {
	return err != nil && err.Error() == fmt.Sprintf(`pq: database "%s" already exists`, dbName)
}
