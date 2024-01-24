package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/peterayobami/ultimatetravelsapiserver/apimodels"
	"github.com/peterayobami/ultimatetravelsapiserver/database"
	"github.com/peterayobami/ultimatetravelsapiserver/datamodels"
	"github.com/peterayobami/ultimatetravelsapiserver/models"
)

func PersistFlightRequest(w http.ResponseWriter, r *http.Request) {

	// Initialize the request data
	var requestCredentials apimodels.FlightRequestCredentials

	// Get the body
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading body request", http.StatusInternalServerError)
		return
	}

	// Unmarshal the request body
	deserializeErr := json.Unmarshal(body, &requestCredentials)

	// If unmarshal failed...
	if err != nil {
		panic("Failed to deserialize request" + deserializeErr.Error())
	}

	// Set the departure date
	var departureDate time.Time

	switch requestCredentials.RouteModel {
	case models.OneWay:
		parsedDate, parseErr := time.Parse("yyyy-MM-dd", requestCredentials.OneWay.DepartureDate)
		if err != nil {
			panic("Failed to parse departure date" + parseErr.Error())
		}
		departureDate = parsedDate
	case models.RoundTrip:
		parsedDate, parseErr := time.Parse("yyyy-MM-dd", requestCredentials.RoundTrip.DepartureDate)
		if err != nil {
			panic("Failed to parse departure date" + parseErr.Error())
		}
		departureDate = parsedDate
	case models.MultiCity:
		parsedDate, parseErr := time.Parse("yyyy-MM-dd", requestCredentials.MultiCity.DepartureDate1)
		if err != nil {
			panic("Failed to parse departure date" + parseErr.Error())
		}
		departureDate = parsedDate
	}

	// Set the data credentials
	flightRequest := datamodels.FlightRequest{SearchCredentials: string(body), DepartureDate: departureDate}

	// Persist the flight request
	database.Db.Create(&flightRequest)

	w.Write([]byte("Flight Request Completed"))
}
