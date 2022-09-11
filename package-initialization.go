package main

import (
	"belajar-golang/database"
	"fmt"

	// Blank identifier
	/*
	 Biasanya dalam golang kalo kita mengimport sebuah package itu terus gak di pake.
	 Maka si golang akan protes dan kalo kita pake auto complete atau prettier dari golang, import an tersebut akan di otomatis ketika kita save file.
	 Jadi kita gunakan blank identifier (_) jika kita ingin mengimport package tapi tidak ingin di panggil atau di gunakan
	*/
	_ "belajar-golang/helper_session2"

	// Imported Alias
	/*
		Ketika kita menggunakan function, variable, dsb di dalam package yang kita panggil. Biasanya kita akan memanggil nama package itu tersebut terlebih
		dahulu baru kita panggil function, variable, dsb yang ada di dalam package tersebut.
		Nah bagaimana jika nama package kita panjang dan kita malas memanggilnya ataupun kita membuat package dengan nama yang sama?.

		Maka kita gunakan imported alias. Tulis nama aliasnya terserah kita sebelum memanggil package yang kita ingin panggil.

		WARNING:
		ketika menggunakan imported alias untuk mengimport sebuah package yang kita panggil tapi tidak di gunakan maka sigolang akan protes atau error.
		Kalo ingin import package tapi gak di gunain pake blank identifier(_) liat diatas 👆.
	*/
	help "belajar-golang/helper"
	// other "belajar-golang/other" //error, karena walaupun sudah di panggil/import tapi tidak di gunakan.
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)

	fmt.Println(help.GetGreeting2())
}
