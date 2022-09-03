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

	testFunction7 := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(testFunction7)

	// variadic parameters
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testFunction8 := sumAll(slice...)
	fmt.Println(testFunction8)

	// Function as variable
	goodbye := getGoodBye
	fmt.Println(goodbye("yadi"))

	// Function as parameter
	filter := spamFilter

	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter("Yadi", filter)

	// Anonymous function
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Yadi", blacklist)
	registerUser("Anjing", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("eko", func(name string) bool {
		return name == "root"
	})

	value := 10
	// Loop
	fmt.Println(factorialLoop(value))
	fmt.Println(factorialRecursive(value))
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

// variadic Function
func sumAll(numbers ...int) (total int) {
	for _, value := range numbers {
		total += value
	}
	return
}

func getGoodBye(data1 string) string {
	return data1
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

type FilterUser func(string) bool

func registerUser(name string, filter FilterUser) {
	if filter(name) {
		fmt.Println(name, "You are blocked!")
	} else {
		fmt.Println("Welcome", name)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
