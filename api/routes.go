package api

import (
	"github.com/go-chi/chi"
	"github.com/tomiok/patients-API/patients/web"
)

func routes(services *patients.PatientHTTPService) *chi.Mux {
	r := chi.NewMux()

	r.Get("/patients", services.GetPatientsHandler)
	r.Post("/patients", services.CreatePatientsHandler)
	r.Get("/patients/{patientID}", services.GetPatientsByIDHandler)

	return r
}
