package api

import (
	"github.com/tomiok/patients-API/models"
	"github.com/tomiok/patients-API/patients"
	"github.com/tomiok/patients-API/storage"
)

type Services struct {
	gtw models.PatientGateway
}

func Start(port string) {
	db := storage.ConnectToDB()
	defer db.Close()
	services := Services{
		gtw: patients.NewPatientGateway(db),
	}
	r := routes(&services)
	server := newServer(port, r)

	server.Start()
}
