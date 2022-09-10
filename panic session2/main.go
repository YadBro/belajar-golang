package main

import "fmt"

func EndApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error:", message)
	}
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
	fmt.Println("Yadi")
}
