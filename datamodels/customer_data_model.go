package datamodels

type Customer struct {
	BaseDataModel      `gorm:"embedded"`
	Email              string
	Phone              string
	CountryDialingCode string
	FirstName          string
	LastName           string

	FlightBookings []FlightBooking
}
