package main

import "fmt"

func main() {
	var halo string = "Halo Dunia" // membuat variable halo dengan nilai atau datanya "Halo Dunia"
	fmt.Println(halo)
	fmt.Println(halo)

	halo = "Halo aku Yadi Apriyadi yang ganteng sekali" // menimpa nilai atau data baru pada variable yang ada
	fmt.Println(halo)

	var nama string   // membuat initial variable
	fmt.Println(nama) // ini akan menghasilkan nilai string kosong

	nama = "Yadi Apriyadi" // menimpa nilai atau data baru pada variable yang ada
	fmt.Println(nama)

	varshortcut := "(varshortcut :=) Ini adalah variable shorcut" // shorcutnya
	fmt.Println(varshortcut)

	var angka int = 10
	fmt.Println(angka)

	var angka2 int
	angka2 = 20
	fmt.Println(angka2)

	angka3 := 30
	fmt.Println(angka3)

	/*
		CONSTANTA VARIABLE merupakan variable yang nilainya tidak dapat di buat atau di timpa kecuali ada nilai yang
		masuk ke constanta variable itu bisa berubah alias bertambah nialinya.
	*/

	const tidakDapatBerubah string = "Aku tidak dapat di ubah" // kita tidak diwajibkan kasih tau tipe datanya kalo pake constant variable (const)
	fmt.Println(tidakDapatBerubah)

	const tidakDapatBerubah2 = "Tanpa di kasih tau tipe datanya" // golangnya nanti dapat mengerti kalo ini string
	fmt.Println(tidakDapatBerubah2)

	/*
	 const tidakDapatBerubah3 string // kita tidak bisa membuat initial variable dengan constanta (const) karena dia harus wajib di masukkan nilai atau harus ada nilainya
	*/

	// Multiple variable
	var (
		nama_depan    = "Yadi"
		nama_belakang = " Apriyadi"
	)

	fmt.Println(nama_depan + nama_belakang)

}
