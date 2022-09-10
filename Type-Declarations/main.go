package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var myKTP noKTP = "12321341214213434"
	var myWife Married = true
	fmt.Println(myKTP)
	fmt.Print(myWife)
}
