package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomiok/patients-API/models"
	"net/http"
	"strconv"
)

func (s *Services) getPatientsHandler(w http.ResponseWriter, r *http.Request) {
	patients := s.GetPatients()
	Success(&patients, http.StatusOK).Send(w)
}

func (s *Services) getPatientsByIDHandler(w http.ResponseWriter, r *http.Request) {
	patientID := chi.URLParam(r, "patientID")
	id, _ := strconv.ParseInt(patientID, 10, 64)
	patient, err := s.GetPatientByID(id)

	if err != nil {
		ErrBadRequest.Send(w)
		return
	}

	Success(&patient, http.StatusOK).Send(w)
}

func (s *Services) createPatientsHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	var cmd models.CreatePatientCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		ErrInvalidJSON.Send(w)
		return
	}

	patient, err := s.CreatePatient(&cmd)

	if err != nil {
		ErrBadRequest.Send(w)
		return
	}

	Success(&patient, http.StatusOK).Send(w)

}
