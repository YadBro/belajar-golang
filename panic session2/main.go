package main

import "fmt"

func EndApp() {
	fmt.Println("Aplikasi Selesai")
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	RunApp(true)
}
