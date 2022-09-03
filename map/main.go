package main

import "fmt"

func main() {

	// Create new map long version
	var people map[string]string = map[string]string{
		"name":    "Yadi Apriyadi",
		"address": "Cimahi",
	}

	// Create new map short version
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
		0: map[string]string{ // ini dry / redundant ikuti dry principles!
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

	// Add or Update new map data
	students[3] = map[string]string{
		"name":    "saipul",
		"address": "malang",
	}

	// Results
	fmt.Println(students)

	// Add or Update new map data
	students[2] = map[string]string{
		"name":    "udin update",
		"address": "jepara update",
	}

	// Results
	fmt.Println(students)

	// How to access
	fmt.Println("name :", students[3]["name"], "\naddress :", students[3]["address"])

	// Read all
	fmt.Print("\n\n=============== ALL ===============\n")
	for key, value := range students {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
		fmt.Println("students[key]: ", students[key])
		fmt.Println("students[key][name]: ", students[key]["name"])
		fmt.Println("students[key][address]: ", students[key]["address"])
		fmt.Println("name :", value["name"], "\naddress :", value["address"], "\n''")
	}

	// Create new map make function version
	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	// Results
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
