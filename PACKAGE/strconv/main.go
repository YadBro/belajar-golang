package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	boolean, error := strconv.ParseBool("true")
	// boolean, error := strconv.ParseBool("ups") // tidak bisa di parsing ke string
	if error == nil {
		fmt.Println("ParseBool :", boolean, reflect.TypeOf(boolean))
	} else {
		fmt.Println(error.Error())
	}

	number, error := strconv.ParseInt("1000000", 10, 64)
	if error == nil {
		fmt.Println("ParseNumber :", number, reflect.TypeOf(number))
	} else {
		fmt.Println(error.Error())
	}

	formatInt := strconv.FormatInt(1000000, 10)
	fmt.Println("FormatInt :", formatInt, reflect.TypeOf(formatInt))

	// Recommended
	// Atoi = string > integer
	// Itoa = integer > string
	valueInt, error := strconv.Atoi("999999999999999999")
	if error == nil {
		fmt.Println("Atoi :", valueInt)
	} else {
		fmt.Println("Error :", error.Error())
	}

	valueString := strconv.Itoa(999999999999999999)
	fmt.Println("Itoa :", valueString)
}
