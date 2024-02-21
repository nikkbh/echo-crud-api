package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	conn := os.Getenv("DB_CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database.")
	}
}
