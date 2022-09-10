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
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(SumAll(slice...))
}
