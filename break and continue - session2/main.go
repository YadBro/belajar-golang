package main

import "fmt"

func main() {
	fmt.Println("BREAK =============")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("\n\nGANJIL ============")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("\n\nGENAP ============")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Perulangan ke", i)
		}
	}
}
