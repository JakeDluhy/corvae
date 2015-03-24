package models

import (
	"time"
	"database/sql"
	"corvae/app/helpers"
	"log"
)

// Database schema
type Patient struct {
	ID int
	Name string
	Age int
	Email string `sql: "not null; unique"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Has one ECG
	EcgSignal EcgStream
	EcgSignalID sql.NullInt64

	// Has many important ecgs
	ImportantEcgs []EcgPiece

	// Has many Doctors: See Doctor model for definitioin

	// Belongs to Nurse Station
	NurseStation NurseStation
	NurseStationID sql.NullInt64
}

func (p *Patient) BeforeCreate() (err error) {
	hashedPass, err := helpers.HashPassword(p.Password)
	if err != nil {
		log.Fatal(err)
	}
	p.Password = hashedPass
	return
}

func (p *Patient) LoginPatient(password string) (err error) {
	return helpers.ComparePasswords(p.Password, password)
}