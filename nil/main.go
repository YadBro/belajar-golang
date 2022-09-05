package main

import (
	"fmt"
)

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person = NewMap("Yadi")
	if person == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person)
	}
}
