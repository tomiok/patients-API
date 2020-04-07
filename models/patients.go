package models

import "time"

type Patient struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePatientCMD struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type PatientGateway interface {
	CreatePatient(p *Patient) (*Patient, error)
	GetPatients() []*Patient
	GetPatientByID(id int64) (*Patient, error)
}

type PatientStorage interface {
	Save(p *Patient) (*Patient, error)
	Get() []*Patient
	GetByID(id int64) (*Patient, error)
}
