package api

import (
	"github.com/tomiok/patients-API/internal/storage"
	patients "github.com/tomiok/patients-API/patients/web"
)

func Start(port string) {
	db := storage.ConnectToDB()
	defer db.Close()

	r := routes(patients.NewPatientHTTPService(db))
	server := newServer(port, r)

	server.Start()
}
