package models

import (
	"time"
	"database/sql"
	// "corvae/db"
)

// Database schema
type Patient struct {
	ID int
	Name string
	Age int
	Email string `sql: "not null; unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	EcgSignal Ecg
	EcgSignalID sql.NullInt64
}