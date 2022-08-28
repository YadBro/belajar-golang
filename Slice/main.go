package main

import (
	"fmt"
	"reflect"
)

func main() {
	type noKTP string
	months := [12]string{ // array
		"January",
		"February",
		"March",
		"April",
		"Mei",
		"Juny",
		"July",
		"Agust",
		"September",
		"October",
		"November",
		"December",
	}

	days := []string{ // slice
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thrusday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	// result
	fmt.Println("Months :", months)
	fmt.Println("Days :", days)

	// Inspect Months and Days
	fmt.Println("Length from Months :", len(months))
	fmt.Println("Length from Days :", len(days))
	fmt.Println("Capacity from Months :", cap(months))
	fmt.Println("Capacity from Days :", cap(days))

	// types
	fmt.Println("Data Type Months :", reflect.TypeOf(months))
	fmt.Println("Data Type Days :", reflect.TypeOf(days))

	// Manipulation 1
	fmt.Println("'\n\n'")
	fmt.Println("================ Manipulation 1 ================")
	sliceMonths1 := months[5:7]
	fmt.Println("sliceMonths1 :", sliceMonths1)

	// Inspect sliceMonths1
	fmt.Println("Length from sliceMonths1 :", len(sliceMonths1))
	fmt.Println("Capacity from sliceMonths1 :", cap(sliceMonths1))

	// // Manipulation 2
	// fmt.Println("================ Manipulation 2 ================")
	// sliceMonths2 := append(sliceMonths1, "")
	// fmt.Println("sliceMonths2 :", sliceMonths2)

	// // Inspect sliceMonths2
	// fmt.Println("Length from sliceMonths2 :", len(sliceMonths2))
	// fmt.Println("Capacity from sliceMonths2 :", cap(sliceMonths2))

	var myKTP noKTP = "12321341214213434"
	fmt.Println("MyKTP :", myKTP)

}
