package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}

	// Operator AND (&) untuk melakukan pointer atau data yang saling berhubungan.
	// Karena secara default golang akan mengcopy
	address2 := &address1
	address5 := &address1

	address2.City = "Cimahi"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address5)

	address3 := address2

	address3.City = "Subang"
	fmt.Println("DATA=========")
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address5)
	*address3 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println("DATA=========")
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address5)

	// Pointer kosongan
	var address4 = new(Address)
	address4.City = "Jakarta"

	address6 := address4
	fmt.Println(address4)

	fmt.Println("DATA==========")
	address6.City = "Wakanda"
	fmt.Println(address4)
	fmt.Println(address6)
}
