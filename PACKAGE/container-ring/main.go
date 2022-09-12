package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// Create new a ring
	data := ring.New(5)

	// Input data to ring with loop
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
