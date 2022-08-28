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

	// Manipulation Months
	fmt.Println("'\n\n'")
	fmt.Println("================ Manipulation Months 1 ================")
	sliceMonths1 := months[5:7]
	fmt.Println("sliceMonths1 :", sliceMonths1)

	// Inspect sliceMonths1
	fmt.Println("Length from sliceMonths1 :", len(sliceMonths1))
	fmt.Println("Capacity from sliceMonths1 :", cap(sliceMonths1))

	fmt.Println("'\n'")
	fmt.Println("================ Manipulation Months 1 ================")
	sliceMonths1 = append(sliceMonths1, "Abang jago")
	fmt.Println("sliceMonths1 :", sliceMonths1)
	fmt.Println("Months :", months)

	// Inspect sliceMonths1
	fmt.Println("Length from sliceMonths1 :", len(sliceMonths1))
	fmt.Println("Capacity from sliceMonths1 :", cap(sliceMonths1))

	// Manipulation Days
	fmt.Println("'\n\n'")
	fmt.Println("================ Manipulation Days 1 ================")
	sliceDays := days[6:7]
	fmt.Println("sliceDays :", sliceDays)

	// Inspect sliceDays
	fmt.Println("Length from sliceDays :", len(sliceDays))
	fmt.Println("Capacity from sliceDays :", cap(sliceDays))

	// Manipulation Days2
	fmt.Println("'\n'")
	fmt.Println("================ Manipulation Days 2 ================")
	sliceDays = append(sliceDays, "New day!")
	fmt.Println("sliceDays :", sliceDays)
	fmt.Println("days :", days)

	fmt.Println("'\n\n'")
	var myKTP noKTP = "12321341214213434"
	fmt.Println("MyKTP :", myKTP)

}
