package model

type Customer struct {
	Firstname string   `json:"firstName"`
	Lastname  string   `json:"lastName"`
	Age       int      `json:"age"`
	Hobbies   []string `json:"hobbies"`
	Addresses []Address
}

type Address struct {
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}
