package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "corvae/app/models"
)

var db, _ = gorm.Open("postgres", "dbname=corvae sslmode=disable")


func Start() {
	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Auto Migrate all models
	db.AutoMigrate(
		&models.Patient{},
		&models.Doctor{},
		&models.NurseStation{},
		&models.EcgStream{},
		&models.EcgPiece{},
	)
}

func GetDB() gorm.DB {
	return db
}