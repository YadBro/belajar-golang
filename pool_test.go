package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return "Default Value"
		},
	}

	pool.Put("Yadi")
	pool.Put("Apriyadi")
	pool.Put("Ganteng")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	println("Selesai")
}
