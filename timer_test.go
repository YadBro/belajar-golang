package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	timeChannel := <-timer.C
	fmt.Println(timeChannel)
}

func TestAfter(t *testing.T) {
	timer := time.After(5 * time.Second)
	fmt.Println(time.Now())

	timeChannel := <-timer
	fmt.Println(timeChannel)
}

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
