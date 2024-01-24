package database

import (
	"log"
	"os"

	"github.com/peterayobami/ultimatetravelsapiserver/datamodels"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Migrate() {
	var err error
	connStr := os.Getenv("DB_CONN")
	Db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database: " + err.Error())
		return
	}

	Db.AutoMigrate(&datamodels.AmadeusCredential{},
		&datamodels.Customer{},
		&datamodels.FlightBooking{},
		&datamodels.FlightOffer{},
		&datamodels.FlightRequest{},
		&datamodels.Payment{},
		&datamodels.Traveler{})
}
