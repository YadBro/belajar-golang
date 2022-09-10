package main

import "fmt"

type Filter func(string) string

func SayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func SpamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func SpamFilter2(name string) string {
	switch {
	case name == "Anjing":
		return "..."
	default:
		return name
	}
}

func SpamFilter3(name string) string {
	switch name {
	case "Anjing":
		return "..."
	default:
		return name
	}
}

func main() {
	filter := SpamFilter
	SayHelloWithFilter("Anjing", filter)
	SayHelloWithFilter("Yadi", SpamFilter2)
	SayHelloWithFilter("Eko", SpamFilter3)
}
