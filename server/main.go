package main

import (
	"github.com/tomiok/patients-API/db"
	"os"

	"github.com/tomiok/patients-API/api"
)

const defaultPort = "8080"

func main() {
	conn := db.ConnectToDB()
	defer conn.Close()
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
