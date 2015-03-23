package models

import (
	"time"
	"database/sql"
)

// Database schema
type Ecg struct {
	ID int
	Signal []sql.NullInt64
	Created_at time.Time
	Updated_at time.Time
}