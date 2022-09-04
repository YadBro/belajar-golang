package main

import "fmt"

// Struct ibarat kata adalah abstract class yang dimana setiap property dan method pada abstract class itu harus ada pada childnya

// Atau mungkin bisa di bilang sebuah object di javascript

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

// Struct function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, "My name is", customer.Name)
}

func (a Customer) sayBudug(name string) {
	fmt.Println("Hi", name, "I am", a.Name, "You budug")
}

func main() {
	var yadi Customer
	yadi.Name = "Yadi"
	yadi.Address = "Cimahi"
	yadi.Age = 19

	fmt.Println(yadi)

	// Struct literals
	joko := Customer{
		Name:    "Joko",
		Address: "Cimahi",
		Age:     20,
	}

	fmt.Println(joko)

	// budi := Customer{"Budi", "Banten", 20} // error
	// fmt.Println(budi)

	yadi.sayHi("Joko")
	joko.sayBudug("Watasi")
}
