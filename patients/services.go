package patients

import (
	"database/sql"
	"github.com/tomiok/patients-API/models"
	"log"
	"time"
)

/**
CreatePatient(p *CreatePatientCMD) (*Patient, error)
GetPatients() []*Patient
GetPatientByID(id int64) (*Patient, error)

    first_name text NOT NULL,
    last_name text NOT NULL,
    address text NOT NULL,
    phone text NOT NULL,
    email text NOT NULL,
*/

type PatientService struct {
	db *sql.DB
}

func (s *PatientService) CreatePatient(p *models.CreatePatientCMD) (*models.Patient, error) {
	res, err := s.db.Exec("insert into patient (first_name, last_name, address, phone, email) values (?,?,?,?,?)",
		p.FirstName, p.LastName, p.Address, p.Phone, p.Email)

	id, err := res.LastInsertId()

	if err != nil {
		log.Println("cannot save the patient")
		return nil, err
	}

	return &models.Patient{
		ID:        id,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Address:   p.Address,
		Phone:     p.Phone,
		Email:     p.Email,
		CreatedAt: time.Now(),
	}, nil
}

func (s *PatientService) GetPatients() []*models.Patient {
	rows, err := s.db.Query("select id, first_name, last_name, address, phone, email, created_at from patient")

	if err != nil {
		log.Printf("cannot execute select query: %s", err.Error())
		return nil
	}
	defer rows.Close()
	var patients []*models.Patient
	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(&patient.ID, &patient.FirstName)
		if err != nil {
			log.Println("cannot read current row")
			return nil
		}
		patients = append(patients, &patient)
	}

	return patients
}
