package datamodels

type Traveler struct {
	BaseDataModel   `gorm:"embedded"`
	FlightBookingId string
	Title           string
	FirstName       string
	LastName        string
	MiddleName      *string
	Gender          string
	TravelerId      string
	TravelerType    string
	DateOfBirth     string

	FlightBooking FlightBooking
}
