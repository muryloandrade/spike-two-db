package database

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_db")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")

	return db
}
