package patients

import (
	"database/sql"
	"github.com/tomiok/patients-API/patients/models"
	"log"
	"time"
)

type PatientStorage interface {
	createPatientDB(p *patients.CreatePatientCMD) (*patients.Patient, error)
	getPatientsDB() []*patients.Patient
	getPatientByIDBD(id int64) (*patients.Patient, error)
}

type PatientService struct {
	db *sql.DB
}

func NewPatientStorageGateway(db *sql.DB) PatientStorage {
	return &PatientService{db: db}
}

func (s *PatientService) createPatientDB(p *patients.CreatePatientCMD) (*patients.Patient, error) {
	log.Println("creating a new patient")
	res, err := s.db.Exec("insert into patient (first_name, last_name, address, phone, email) values (?,?,?,?,?)",
		p.FirstName, p.LastName, p.Address, p.Phone, p.Email)

	if err != nil {
		log.Printf("cannot save the patient, %s", err.Error())
		return nil, err
	}

	id, err := res.LastInsertId()

	return &patients.Patient{
		ID:        id,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Address:   p.Address,
		Phone:     p.Phone,
		Email:     p.Email,
		CreatedAt: time.Now(),
	}, nil
}

func (s *PatientService) getPatientsDB() []*patients.Patient {
	rows, err := s.db.Query("select id, first_name, last_name, address, phone, email, created_at from patient")

	if err != nil {
		log.Printf("cannot execute select query: %s", err.Error())
		return nil
	}
	defer rows.Close()
	var p []*patients.Patient
	for rows.Next() {
		var patient patients.Patient
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone,
			&patient.Email, &patient.CreatedAt)
		if err != nil {
			log.Println("cannot read current row")
			return nil
		}
		p = append(p, &patient)
	}

	return p
}

func (s *PatientService) getPatientByIDBD(id int64) (*patients.Patient, error) {
	var patient patients.Patient
	err := s.db.QueryRow(`select id, first_name, last_name, address, phone, email, created_at from patient
		where id = ?`, id).Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone,
		&patient.Email, &patient.CreatedAt)

	if err != nil {
		log.Printf("cannot fetch patient")
		return nil, err
	}

	return &patient, nil
}
