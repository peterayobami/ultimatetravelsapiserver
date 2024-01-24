package datamodels

type FlightOffer struct {
	BaseDataModel `gorm:"embedded"`
	AmaClientRef  string
	Data          string
}
