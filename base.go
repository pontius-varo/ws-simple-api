package main

type Reservation struct {
	MainFirstName     string   `json:"mainFirstName"`
	MainLastName      string   `json:"mainLastName"`
	AdditionalPersons []Person `json:"additionalPersons"`
}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
