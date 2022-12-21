package model

type Customer struct {
	CustomerId string `json:"id"`
	Firstname  string `json:"firstname,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
}
