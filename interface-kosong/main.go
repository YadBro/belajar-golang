package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return i
	} else {
		return "Ups"
	}
}

func Ups2(i interface{}) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return i
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)

	var data2 interface{} = Ups2("1")
	fmt.Println(data2)

	var data3 interface{}
	data3 = "Hello"
	fmt.Println(data3)
	data3 = false
	fmt.Println(data3)
	data3 = 1
	fmt.Println(data3)
}
