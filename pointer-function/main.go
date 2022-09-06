package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Man struct {
	Name string
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
	// return address.Country
}

func ChangeCountryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	var newAddress = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(newAddress)
	fmt.Println(newAddress)

	newAddressPointer := &newAddress

	ChangeCountryToIndonesiaPointer(newAddressPointer)
	fmt.Println(newAddress)
	fmt.Println(newAddressPointer)

	yadi := Man{
		Name: "Yadi",
	}
	yadi.Married()    // Mr. Yadi
	fmt.Println(yadi) // Yadi (tidak berubah)

	udin := Man{
		Name: "Udin",
	}
	udin.MarriedPointer() // Mr. Udin
	fmt.Println(udin)     // Mr. Udin (berubah)
}
