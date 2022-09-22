package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database string
var DSN = "host=localhost user=matias password=mypassword dbname=gorm port=5432"
var DB *gorm.DB

func DbConnection() {
	// esto devuelve la conexion o error
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB conectada")
	}

}
