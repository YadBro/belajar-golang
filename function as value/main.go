package main

import "fmt"

func main() {
	sayGoodBye := GetGoodBye

	result := sayGoodBye("Yadi")
	fmt.Println(result)
}

func GetGoodBye(name string) string {
	return "Good bye " + name
}
