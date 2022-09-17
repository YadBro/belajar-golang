package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yadi")

	if result != "Hello Yadi" {
		// error
		panic("Result is not Hello Yadi")
	}
}
