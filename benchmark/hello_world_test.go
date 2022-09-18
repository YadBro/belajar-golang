package benchmark

import (
	"fmt"
	"testing"
)

type BenchmarkHelloWorld []struct {
	Name    string
	Request string
}

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

// SubBenchmark

func BenchmarkSub(b *testing.B) {
	b.Run("Yadi", func(b *testing.B) {
		for i := 0; i > b.N; i++ {
			HelloWorld("Yadi")
		}
	})

	b.Run("Apriyadi", func(b *testing.B) {
		for i := 0; i > b.N; i++ {
			HelloWorld("Apriyadi")
		}
	})
}

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := BenchmarkHelloWorld{
		{
			Name:    "Yadi",
			Request: "Yadi",
		},
		{
			Name:    "Apriyadi",
			Request: "Apriyadi",
		},
		{
			Name:    "Eko",
			Request: "Eko",
		},
		{
			Name:    "Khannedy",
			Request: "Khannedy",
		},
		{
			Name:    "Kurniawan",
			Request: "Kurniawan",
		},
		{
			Name:    "Joko",
			Request: "Joko",
		},
		{
			Name:    "Budi",
			Request: "Budi",
		},
	}

	for _, becnhmark := range benchmarks {
		b.Run(becnhmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(becnhmark.Request)
			}
		})
	}
}
