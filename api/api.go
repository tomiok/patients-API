package api

import "github.com/tomiok/patients-API/models"

type Services struct {
	models.PatientGateway
}

func Start(port string) {
	services := Services{}
	r := routes(&services)
	server := newServer(port, r)

	server.Start()
}
