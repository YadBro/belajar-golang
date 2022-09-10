package main

import "fmt"

func SayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func SpamFilter(name string) string {
	switch {
	case name == "Anjing":
		return "..."
	default:
		return name
	}
}

func main() {
	filter := SpamFilter
	SayHelloWithFilter("Anjing", filter)
	SayHelloWithFilter("Yadi", SpamFilter)
}
