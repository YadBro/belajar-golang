package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHello(name string) {
	fmt.Println("Hello " + name + ", my Name is " + customer.Name)
}

func (a Customer) SayGoodBye(name string) {
	fmt.Println(name + ": \"Good bye " + a.Name + "\"")
}

func main() {
	var yadi Customer
	yadi.Name = "Yadi Apriyadi"
	yadi.Address = "Jl. Kebon Jeruk"
	yadi.Age = 19

	fmt.Println(yadi)
	fmt.Println(yadi.Name)
	fmt.Println(yadi.Address)
	fmt.Println(yadi.Age)

	// Struct Literals
	joko := Customer{
		Name:    "Joko",
		Address: "Jl. Kenangan Mantan",
		Age:     15,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 5}
	fmt.Println(budi)

	// Struct Method
	yadi.SayHello(joko.Name)
	yadi.SayGoodBye(budi.Name)
}
