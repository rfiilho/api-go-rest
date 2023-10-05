package main

import (
	"fmt"

	"github.com/rfiilho/api-go-rest/database"
	"github.com/rfiilho/api-go-rest/routes"
)

func main() {
	database.Connect()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
