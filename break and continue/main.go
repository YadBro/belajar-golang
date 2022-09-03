package main

import "fmt"

func main() {
	// BREAK
	for i := 0; i <= 10; i++ {
		fmt.Println("This is loop: ", i)
		if i == 5 {
			break
		}
		fmt.Println("This is loop: ", i)
	}

	//  CONTINUE
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			continue
		}
		fmt.Println("This is loop: ", i)
	}
}
