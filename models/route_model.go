package models

type RouteModel int

const (
	OneWay    RouteModel = iota // 1
	RoundTrip                   // 2
	MultiCity                   // 3
)
