package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import necessário para o driver PostgreSQL
)

var (
	DBONE *sql.DB
	DBTWO *sql.DB
)

func ConnectDB() {
	// Read database connection parameters from environment variables or configuration
	host := "172.20.120.180" // e.g., "localhost"
	port := "5433"           // e.g., "5432"
	portTwo := "5432"        // e.g., "5432"
	dbname := "root"         // e.g., "users"
	user := "root"           // e.g., "root"
	password := "root"       // e.g., "root"

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	dsnTwo := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, portTwo, dbname, user, password)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatal("Error testing the database connection:", err)
	}

	DBONE = db

	fmt.Println("Successfully connected!")

	dbTwo, err := sql.Open("postgres", dsnTwo)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Test the connection
	if err = dbTwo.Ping(); err != nil {
		log.Fatal("Error testing the database connection:", err)
	}

	DBTWO = dbTwo

	fmt.Println("Successfully connected two!")

}
