package main

import (
	"log"
	config "my-trips-api/configs"
	"my-trips-api/internal/database"
	"my-trips-api/internal/routes"
)

func main() {

	_, err := config.LoadConfig("cmd/server")
	if err != nil {
		log.Panic("Error loading config")
	}

	database.ConnectionToDatabase()
	routes.HandlerRequests()

}
