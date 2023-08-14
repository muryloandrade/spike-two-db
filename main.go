package main

import (
	"connect-db/api/database"
	"connect-db/api/routes"
	"fmt"
)

func main() {

	database.ConnectDB()

	fmt.Println("Iniciando servidor, redirecionando para home...")

	routes.HandleResquests()
}
