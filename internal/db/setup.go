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

	_, err = defaultDB.Exec("CREATE DATABASE " + dbName)
	if err != nil && !isDuplicateDBError(err) {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Connect to the newly created DB
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to new DB: %v", err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS inventory")
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS bins")
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE inventory (
		id SERIAL PRIMARY KEY,
		item_name TEXT UNIQUE,
		bin_id TEXT,
		description TEXT
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE bins (
		id SERIAL PRIMARY KEY,
		bin_name TEXT UNIQUE,
		description TEXT DEFAULT ''
	)`)
	if err != nil {
		log.Fatalf("Failed to create bins: %v", err)
	}

	// Insert sample data
	_, err = db.Exec(`INSERT INTO inventory (item_name, bin_id) VALUES 
		('Screwdriver', 1), 
		('Hammer', 2), 
		('Wrench', 1)`)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	_, err = db.Exec(`INSERT INTO bins (bin_name) VALUES 
		('Costumes'), 
		('Props'), 
		('Weapons')`)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	log.Println("Database, table, and sample data set up successfully.")
	return db
}

func isDuplicateDBError(err error) bool {
	return err != nil && err.Error() == fmt.Sprintf(`pq: database "%s" already exists`, dbName)
}
