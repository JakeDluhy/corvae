package models

import (
	"time"
)

// Database schema
type NurseStation struct {
	ID int
	StationID string
	StationPassword string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Has many Patients
	Patients []Patient
}