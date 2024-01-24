package datamodels

import "time"

type FlightRequest struct {
	BaseDataModel     `gorm:"embedded"`
	CustomerId        *string
	SearchCredentials string
	DepartureDate     time.Time
}
