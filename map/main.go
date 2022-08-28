package main

import "fmt"

func main() {
	var people map[string]string = map[string]string{
		"name":    "Yadi Apriyadi",
		"address": "Cimahi",
	}

	person := map[int]string{
		0: "yadi",
		1: "bambang",
		2: "udin",
	}

	// Results
	fmt.Println(person)
	fmt.Println(people)

	// Access map
	fmt.Println(person[0])
	fmt.Println(people["name"])

	students := map[int]map[string]string{
		0: map[string]string{ // versi panjang
			"name":    "yadi",
			"address": "cimahi",
		},
		1: { // mending gini aja
			"name":    "bambang",
			"address": "subang",
		},
		2: {
			"name":    "udin",
			"address": "jepara",
		},
	}

	// Results
	fmt.Println(students)

	// How to access
	fmt.Println("name :", students[0]["name"], "\naddress :", students[0]["address"])
}
