// Package container / list adalah implementasi struktur data double linked list di Go-lang.

package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Yadi")
	data.PushBack("Apriyadi")
	data.PushBack("Ganteng")

	fmt.Println(data)

	fmt.Println(data.Front().Value)  // mengambil data awal
	fmt.Println(data.Back().Value)   // mengambil data akhir
	fmt.Println(data.Front().Prev()) // mengambil data sebelumnya setelah data awal
	fmt.Println(data.Back().Next())  // mengambil data selanjutnya setelah data akhir

	data.PushFront("Budi")
	fmt.Println(data.Front().Value) // mengambil data awal
	fmt.Println(data.Back().Value)  // mengambil data akhir

	fmt.Println("\n\nAmbil semua dari depan ============")
	// NEXT
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("\n\nAmbil semua dari belakang ============")
	// REVERSE
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
