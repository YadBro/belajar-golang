package main

import "fmt"

func Logging() {
	fmt.Println("Selesai memanggil longging!")
}

func RunApplication(num int) {
	defer Logging()
	fmt.Println("Run Application")
	result := 10 / num
	fmt.Println("Result", result)
}

func main() {
	RunApplication(0)
}
