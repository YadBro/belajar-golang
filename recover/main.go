package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
	errorMessage := recover()
	fmt.Println("Error: ", errorMessage)
}

func runApplication(error bool) {
	defer logging()
	defer func() {
		fmt.Println("KAMU PANIC YA")
	}()
	if error {
		panic("Application Error!")
	}
	fmt.Println("Run Application")
}

func main() {
	// Recofer simplenya seperti block catch yang ada di javascript yaitu ketika ada error program tetap jalan tetapi pesan errornya akan di tangkap.
	// Ibarat kata seperti try catch, defer ini adalah try nya dan catch adalah recover

	runApplication(false) // success
	runApplication(true)  // error
}
