package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hi", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	yadi := Person{
		Name: "Yadi",
	}

	cat := Animal{
		Name: "Meow",
	}

	sayHello(yadi)
	sayHello(cat)

	var data interface{} = Ups(3)
	fmt.Println(data)
}
