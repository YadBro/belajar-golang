package main

import "fmt"

func main() {
	type noKTP string

	var myKTP noKTP = "12321341214213434"
	fmt.Println(myKTP)
}
