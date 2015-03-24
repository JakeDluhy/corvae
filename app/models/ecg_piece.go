package models

import (
	"time"
	"database/sql"
)

// Database schema
type EcgPiece struct {
	ID int
	Signal []sql.NullInt64
	StartTime time.Time
	Created_at time.Time
	Updated_at time.Time
}