package datamodels

type Payment struct {
	BaseDataModel   `gorm:"embedded"`
	Reference       string
	FlightBookingId string
	PaymentStatus   string

	FlightBooking FlightBooking
}
