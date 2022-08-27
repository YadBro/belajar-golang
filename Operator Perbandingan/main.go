package main

import "fmt"

func main() {
	a := 1
	b := 1
	c := 2

	result1 := a == b // true
	fmt.Println(result1)

	result2 := a != b // false
	fmt.Println(result2)

	result3 := a < c // true
	fmt.Println(result3)

	result4 := c < a // false
	fmt.Println(result4)

	result5 := a > c // false
	fmt.Println(result5)

	result6 := c > a // true
	fmt.Println(result6)

	result8 := a <= b // true
	fmt.Println(result8)

	result7 := a <= c // true
	fmt.Println(result7)

	result9 := c <= a // false
	fmt.Println(result9)

	result10 := a >= b // true
	fmt.Println(result10)

	result11 := a >= c // false
	fmt.Println(result11)

	result12 := c >= a // true
	fmt.Println(result12)
}
