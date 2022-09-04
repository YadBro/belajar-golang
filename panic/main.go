package main

import "fmt"

func endApp() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(err bool) {
	defer endApp()
	if err {
		panic("ERROR")
	}
	fmt.Println("Run Application")
}

func main() {
	// Ketika ada error sebenarnya panic function di jalankan, fungsinya panic function adalah untuk menghentikan program ketika adanya sebuah error.

	runApplication(false) // success
	runApplication(true)  // error
}
