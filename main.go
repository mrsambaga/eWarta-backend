package main

import (
	"log"
	"stage01-project-backend/db"
	"stage01-project-backend/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
		return
	}

	server.Init()
}
