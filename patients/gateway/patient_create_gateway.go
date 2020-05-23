package patients

import (
	"database/sql"
	"github.com/tomiok/patients-API/patients/models"
)

type PatientGateway interface {
	CreatePatient(p *patients.CreatePatientCMD) (*patients.Patient, error)
	GetPatients() []*patients.Patient
	GetPatientByID(id int64) (*patients.Patient, error)
}
type CreatePatientInDB struct {
	PatientStorage
}

func (c *CreatePatientInDB) CreatePatient(p *patients.CreatePatientCMD) (*patients.Patient, error) {
	return c.createPatientDB(p)
}

func (c *CreatePatientInDB) GetPatients() []*patients.Patient {
	return c.GetPatients()
}

func (c *CreatePatientInDB) GetPatientByID(id int64) (*patients.Patient, error) {
	return c.getPatientByIDBD(id)
}

func NewPatientGateway(db *sql.DB) PatientGateway {
	return &CreatePatientInDB{NewPatientStorageGateway(db)}
}
