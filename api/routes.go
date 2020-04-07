package api

import (
	"github.com/go-chi/chi"
)

func routes(services *Services) *chi.Mux {
	r := chi.NewMux()

	r.Get("/patients", services.getPatientsHandler)
	r.Post("/patients", services.createPatientsHandler)
	r.Get("/patients/{patientID}", services.getPatientsByIDHandler)

	return r
}
