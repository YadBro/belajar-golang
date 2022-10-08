package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // Memberikan informasi dimana hasil build aplikasi ini di build jadi aplikasi exe. Beserta argument argumentnya
	/*
		Contohnya:
		go run main.go Yadi Apriyadi

		maka hasil variable args ini adalah:
		[C:\Users\Yadi-PC\AppData\Local\Temp\go-build2662322810\b001\exe\main.exe Yadi Apriyadi]

		kalo tidak pake argument pada saat menjalakan aplikasinya di terminal maka hasilnya seperti ini.
		[C:\Users\Yadi-PC\AppData\Local\Temp\go-build2662322810\b001\exe\main.exe]
	*/
	// ARGS
	fmt.Println("\nArguments :", args)

	// Menampilkan argument-argumentnya
	// Case Study: go run main.go Yadi Apriyadi
	// Noted: Ketika ingin menampilkan argumentnya maka wajib memasukkan berapa argumentnya gak boleh kurang karena kita akan menampilkannya.
	// fmt.Println(args[0]) // Tempat/Path hasil build aplikasi ini tersimpan
	// fmt.Println(args[1]) // Yadi
	// fmt.Println(args[2]) // Apriyadi

	// HOSTNAME
	name, error := os.Hostname()
	if error == nil {
		fmt.Println("\nhostName :", name) // hostName : DESKTOP-QKGLTUF
	} else {
		fmt.Println("error :", error.Error())
	}

	// version := os.Getenv("VERSION") // ini bukan mengambil environment dari aplikasi kita tapi dari os kita.
	os_dekstop := os.Getenv("OS")
	fmt.Println(os_dekstop)
}
