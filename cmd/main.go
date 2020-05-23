package main

import (
	"log"
	"os"

	"github.com/tomiok/patients-API/api"
)

const defaultPort = "8080"

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
