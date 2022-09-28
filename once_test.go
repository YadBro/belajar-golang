package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	// var once sync.Once = sync.Once{}
	var group sync.WaitGroup = sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			OnlyOnce()
		}()
	}

	fmt.Println("Counter", counter)

}
