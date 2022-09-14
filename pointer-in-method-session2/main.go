package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() interface{} {
	man.Name = "Mr. " + man.Name
	return man
}

func main() {
	yadi := Man{"Yadi"}
	yadi.Married()

	fmt.Println(yadi)
}
