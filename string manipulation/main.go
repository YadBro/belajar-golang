package main

import "fmt"

func main() {
	// menggabungkan 1 atau lebih string
	fmt.Println("Belajar" + " Golang")

	// menghitung jumlah karakter atau kata pada string
	fmt.Println("totalnya karakter/kata diatas adalah : ", len("Belajar Golang"))

	// mengambil huruf pertama, tetapi nanti hasilnya berupa byte. contoh ini akan mendapatkan huruf B dan byte dari B adalah 66
	fmt.Println("Belajar Golang"[0]) // hasilnya 66 (byte). B

	// mengambil huruf B dengan awalnya 0 dan A dengan akhirnya 4
	fmt.Println("Belajar Golang"[0:4]) // hasilnya Bela, kalo ini tidak dalam bentuk byte

	fmt.Println("Belajar Golang"[:7])
}
