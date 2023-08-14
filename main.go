package main

import (
	"connect-db/api/routes"
	"fmt"
)

func main() {

	fmt.Println("Iniciando servidor, redirecionando para home...")

	routes.HandleResquests()
}
