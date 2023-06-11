package main

import (
	"github.com/Milkado/api-ars-arcanum/database"
	"github.com/Milkado/api-ars-arcanum/routes"
)

func main() {
	database.ConnectDb()
	routes.HandleRequests()
}
