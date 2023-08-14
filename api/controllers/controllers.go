package controllers

import (
	"connect-db/api/database"
	"connect-db/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func GetDatabaseOne(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()

	rows, err := database.DBONE.QueryContext(r.Context(), "SELECT * FROM users")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	var users []entity.User // Assuming you have defined a User struct to represent the data

	for rows.Next() {
		var user entity.User

		err := rows.Scan(&user.Id, &user.Nome, &user.Cargo) // Scan column values into struct fields
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error after iterating through rows:", err)
	}

	json.NewEncoder(w).Encode(users)
}
