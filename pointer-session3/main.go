package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	fmt.Println("address1", address1)
	var address2 *Address = &address1
	address3 := &address1

	address2.City = "Bandung"
	fmt.Println("address1", address1)
	fmt.Println("address2", address2)
	fmt.Println("address3", address3)

	*address2 = Address{"Malang", "Jawa Barat", "Indonesia"}
	fmt.Println("address1", address1)
	fmt.Println("address2", address2)
	fmt.Println("address3", address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	address5 := address4
	address5.City = "Bandung"
	address5.Province = "Jawa Barat"
	fmt.Println(address4)
	fmt.Println(address5)
}
