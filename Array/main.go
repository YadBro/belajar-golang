package main

import "fmt"

func main() {
	// some ways for create an array.
	var myArray = [3]string{
		"udin",
		"bambang",
		"asep",
		// "saep", // ini error karena melebihi dari kapasitas arranya
	}

	var myBag [2]string
	myBag[0] = "pencil"
	myBag[1] = "bolpoin"

	myNumber := [3]int{
		1,
		2,
		3,
	}

	fmt.Println(myArray)
	fmt.Println(myArray[1])
	fmt.Println(myBag)
	fmt.Println(myBag[0])
	fmt.Println(myNumber)
	fmt.Println(myNumber[2])
}
