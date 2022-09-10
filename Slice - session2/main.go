package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"Mei",
		"Juny",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println("length of slice1 months", len(slice1))
	fmt.Println("capacity of slice1 months", cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei lagi"
	// fmt.Println(months)

	slice2 := months[11:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Other")
	fmt.Println(slice3)

	slice3[1] = "Not a december"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Yadi"
	newSlice[1] = "Apriyadi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
