package main

import "fmt"

func main() {
	// NUMBERS ===============
	fmt.Println("Ini adalah tipe data number")
	// integer
	fmt.Println(1)

	// integer8 -128 ~ 127
	fmt.Println(int8(127))

	// integer16 -32768 ~ 32767
	fmt.Println(int16(32767))

	// integer32 -2147483648 ~ 2147483647
	fmt.Println(int32(2147483647))

	// integer64 -9223372036854775808  ~ 9223372036854775807
	fmt.Println(int64(9223372036854775807))

	// unsignedInteger8 0 ~ 127
	fmt.Println(uint8(127))

	// unsignedInteger16 0 ~ 32767
	fmt.Println(uint16(32767))

	// unsignedInteger32 0 ~ 2147483647
	fmt.Println(uint32(2147483647))

	// unsignedInteger64 0 ~ 9223372036854775807
	fmt.Println(uint64(9223372036854775807))

	// floating point number (decimal)
	fmt.Println(1.2)
	// NUMBERS ===============

	// STRING ============
	fmt.Println("Ini adalah tipe data string")
	fmt.Println("Hello World")
	// STRING ============

	// BOOLEAN ==========
	fmt.Println("Ini adalah tipe data boolean")
	fmt.Println(true, false)
	// BOOLEAN ==========
}
