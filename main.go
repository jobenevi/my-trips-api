package main

import (
	"my-trips-api/database"
	"my-trips-api/routes"
)

func main() {
	database.ConnectionToDatabase()
	routes.HandlerRequests()
}
