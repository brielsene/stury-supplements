package main

import (
	"stury-supplements/database"
	"stury-supplements/routes"
)

func main() {
	database.ConectaComDB()
	routes.HandleRequests()

}
