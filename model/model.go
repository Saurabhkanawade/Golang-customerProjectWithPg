package model

type Customer struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}
