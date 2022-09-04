package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(val int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / val
	fmt.Println(result)
}

func main() {
	// Defer adalah sebuah keyword yang di gunakan untuk menjalankan sebuah function setelah function parent defer tersebut sudah dijalankan walaupun terjadi error sekalipun defer function tersebut akan tetap di jalankan.

	runApplication(10) // success
	runApplication(0)  // error
}
