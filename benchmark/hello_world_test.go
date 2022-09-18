package benchmark

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// BEFORE
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// AFTER
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yadi")

	if result != "Hello Yadi" {
		// error
		panic("Result is not Hello Yadi")
	}
}

func BenchmarkHelloWorldYadi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Yadi")
	}
}
func BenchmarkHelloWorldApriyadi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Apriyadi")
	}
}
