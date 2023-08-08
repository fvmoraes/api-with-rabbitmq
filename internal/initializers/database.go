package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func StartDatabaseConnect() {
	connectionString := "host=172.34.0.2 user=foobar password=foobar dbname=foobar port=5432"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Database Connection Error", err)
	}

}
