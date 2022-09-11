package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Devider(number int, devide int) (int, error) {
	if devide == 0 {
		return number, errors.New("Devide number does not " + strconv.Itoa(number))
	} else {
		return number, nil
	}
}

func main() {
	result, error := Devider(100, 0)
	if error == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println(error.Error())
	}
}
