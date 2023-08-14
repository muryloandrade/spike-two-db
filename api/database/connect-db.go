package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import necessário para o driver PostgreSQL
)

var db *sql.DB

func ConnectDB() *sql.DB {
	var (
		host     = "localhost"
		port     = 5433
		dbname   = "root"
		user     = "root"
		password = "root"
	)

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	db, err := sql.Open("postgres", dsn)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")

	return db
}
