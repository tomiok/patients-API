package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectToDB() *sql.DB {
	server := os.Getenv("GCP_HOSTNAME")
	if server == "" {
		server = "localhost"
	}
	connURL := fmt.Sprintf("postgres://postgres:postgres@%s/project?sslmode=disable", server)
	db, err := sql.Open("postgres", connURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB via %s: %v", connURL, err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB via %s: %v", connURL, err)
	}
	log.Println("Connected to DB")
	return db
}
