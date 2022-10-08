package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
	Age  int    `required:"true" max:"10"` // struct tag
}

type Sample2 struct {
	Name string
	Age  int
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			} else if value == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	sample := Sample{
		Name: "Yadi",
		Age:  15,
	}

	var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(1).Tag.Get("required"))
	fmt.Println(sampleType.Field(1).Tag.Get("max"))

	sample2 := Sample2{"Bro", 13}
	sample2.Name = ""
	fmt.Println(IsValid(sample))
	fmt.Println(sample)
	fmt.Println(IsValid(sample2))

	recursive(20)
}
func recursive(v int) {
	if v > 21 {
		fmt.Println(true)
	} else {
		fmt.Println(v)
		recursive(v + 1)
	}
}
