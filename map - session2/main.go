package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yadi",
		"address": "Cimahi",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko Kurniawan Khannedy"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
