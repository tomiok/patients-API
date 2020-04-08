package patients

import (
	"database/sql"
	"github.com/tomiok/patients-API/models"
	"log"
	"time"
)

type PatientService struct {
	db *sql.DB
}

func (s *PatientService) CreatePatient(p *models.CreatePatientCMD) (*models.Patient, error) {
	res, err := s.db.Exec("insert into patient (first_name, last_name, address, phone, email) values (?,?,?,?,?)",
		p.FirstName, p.LastName, p.Address, p.Phone, p.Email)

	if err != nil {
		log.Println("cannot save the patient")
		return nil, err
	}

	id, err := res.LastInsertId()

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
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone,
			&patient.Email, &patient.CreatedAt)
		if err != nil {
			log.Println("cannot read current row")
			return nil
		}
		patients = append(patients, &patient)
	}

	return patients
}

func (s *PatientService) GetPatientByID(id int64) (*models.Patient, error) {
	var patient models.Patient
	err := s.db.QueryRow(`select id, first_name, last_name, address, phone, email, created_at from patient
		where id = ?`, id).Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone,
		&patient.Email, &patient.CreatedAt)

	if err != nil {
		log.Printf("cannot fetch patient")
		return nil, err
	}

	return &patient, nil
}
