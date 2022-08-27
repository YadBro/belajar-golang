package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x float32 = 3.2
	tes := int32(x)
	fmt.Println(tes)

	var nilai string = "100"
	nilaiInt, _ := strconv.Atoi(nilai)
	fmt.Println(nilaiInt)

	nilaiString := strconv.Itoa(-100)
	fmt.Println(nilaiString)

	nilaiInt2, _ := strconv.ParseInt("100", 0, 32)
	fmt.Println(nilaiInt2)

	string1 := strconv.FormatInt(42, 0)
	fmt.Println(string1)
}
