package main

import (
	"github.com/tomiok/patients-API/db"
	"log"
	"os"

	"github.com/tomiok/patients-API/api"
)

const defaultPort = "8080"

func main() {
	log.Println("stating API server")
	conn := db.ConnectToDB()
	defer conn.Close()
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
