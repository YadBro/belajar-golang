package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address Address) {
	address.City = "Indonesia"
}

func ChangeCountryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{
		City:     "Sukabumi",
		Province: "Jawa Barat",
		Country:  "Belanda",
	}

	address2 := new(Address)
	address2.City = "Banten"
	address2.Province = "Jawa Barat"
	address2.Country = "Belanda"

	ChangeCountryToIndonesia(address1)
	fmt.Println(address1)

	ChangeCountryToIndonesiaPointer(address2)
	fmt.Println(address2)
}
