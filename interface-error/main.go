package main

import (
	"errors"
	"fmt"
)

func Devide(value int, devide int) (int, error) {
	if devide == 0 {
		return 0, errors.New("Ups devide does not 0")
	} else {
		results := value / devide
		return results, nil
	}
}

func main() {
	result, err := Devide(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err)
	}
}
