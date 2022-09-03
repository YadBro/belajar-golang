package main

import (
	"fmt"
)

func main() {
	testFunction1 := sum([]int{1, 2, 3, 4, 5, 6, 8, 9, 0, 0})
	fmt.Println(testFunction1)

	testFunction2, data1, data2 := fullName("Yadi", "Apriyadi")
	fmt.Println(testFunction2, data1, data2)

	testFunction3 := name("Apriyadi")
	fmt.Println(testFunction3)

	testFunction4 := name2("Yadi")
	fmt.Println(testFunction4)

	testFunction5, _ := name3("Bambang")
	fmt.Println(testFunction5)

	testFunction6, _ := name4("Bambang")
	fmt.Println(testFunction6)

}

// single return
func sum(data []int) (data2 int) {
	for _, value := range data {
		data2 += value
	}
	return data2
}

// multiple return
func fullName(name1 string, name2 string) (string, string, string) {
	data := name1 + name2
	return data, name1, name2
}

// return named
func name(name1 string) (name string) {
	name = name1
	return name
}

// Just return
func name2(name1 string) (name string) {
	name = name1
	return name
}

// Ignore return
func name3(name1 string) (name string, name2 string) {
	name = name1
	return name, name2
}

// return
func name4(name1 string) (name string, name2 string) {
	name = name1
	return
}
