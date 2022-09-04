package main

import "fmt"

func main() {
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println(counter)
	fmt.Println(increment())
}
