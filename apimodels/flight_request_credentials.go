package apimodels

import (
	"github.com/peterayobami/ultimatetravelsapiserver/models"
)

type FlightRequestCredentials struct {
	CurrencyCode  string
	DirectFlight  bool
	FlexibleDates bool
	RouteModel    models.RouteModel
	FlightCabin   string
	Travelers     TravelersSpecification
	OneWay        OneWayRouteModel
	RoundTrip     RoundTripRouteModel
	MultiCity     MultiCityRouteModel
}

type TravelersSpecification struct {
	NumberOfAdults   int
	NumberOfChildren int
	NumberOfInfants  int
}

type OneWayRouteModel struct {
	OriginLocation    string
	OriginDestination string
	DepartureDate     string
}

type RoundTripRouteModel struct {
	OneWayRouteModel
	ReturnDate string
}

type MultiCityRouteModel struct {
	OriginLocation1      string
	DestinationLocation1 string
	DepartureDate1       string
	OriginLocation2      string
	DestinationLocation2 string
	DepartureDate2       string
	OriginLocation3      string
	DestinationLocation3 string
	DepartureDate3       string
	OriginLocation4      string
	DestinationLocation4 string
	DepartureDate4       string
	OriginLocation5      string
	DestinationLocation5 string
	DepartureDate5       string
}
