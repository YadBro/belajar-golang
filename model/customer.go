package model

type Customer struct {
	Firstname string   `json:"firstName"`
	Lastname  string   `json:"lastName"`
	Age       int      `json:"age"`
	Hobbies   []string `json:"hobbies"`
}
