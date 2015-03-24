package models

import (
	"time"
)

// Database schema
type Doctor struct {
	ID int
	Name string
	Email string `sql: "not null; unique"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time

	Patients []Patient `gorm:"many2many:doctor_patients;"`
}