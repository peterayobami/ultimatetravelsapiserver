package datamodels

type FlightBooking struct {
	BaseDataModel `gorm:"embedded"`
	CustomerId    string
	Pnr           *string
	OrderId       *string
	AmaClientRef  string
	BookingStatus string
	PaymentStatus string

	Customer  Customer
	Travelers []Traveler
}
