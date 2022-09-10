package main

import "fmt"

func SumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	fmt.Println(SumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 0))
}
