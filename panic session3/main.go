package main

import "fmt"

func endApp() {
	// message := recover()
	// if message != nil {
	// 	fmt.Println("Terjadi error:", message)
	// }
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error!")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("yadi")
}
