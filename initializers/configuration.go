package initializers

import (
	"log"

	"github.com/lpernett/godotenv"
)

func LoadConfiguration() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading configuration")
	}
}
