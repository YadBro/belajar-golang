package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() interface{} {
	man.Name = "Mr. " + man.Name
	return man
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	yadi := Man{
		Name: "Yadi",
	}

	yadi.Married()
	fmt.Println(yadi.Name)

	joko := Man{
		Name: "Joko",
	}

	joko.MarriedPointer()
	fmt.Println(joko.Name)
}
