package main

import "fmt"

func main() {
	var result = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	case bool:
		fmt.Println("Value", value, "is boolean")
	default:
		fmt.Println("Unknown type")
	}
}

func random() interface{} {
	return nil
}
