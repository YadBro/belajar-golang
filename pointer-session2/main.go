package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Cimahi", "Jawa Barat", "Indonesia"}
	address2 := &address1
	addresstest := address1

	fmt.Println("=========== Single pointer 1")
	address2.City = "Bandung"
	fmt.Println(address1)    // berubah
	fmt.Println(address2)    // berubah
	fmt.Println(addresstest) // tidak berubah

	fmt.Println("\n\n=========== Single pointer 2 with override address2")
	address2 = &Address{"Malang", "Jawa Barat", "Indonesia"}
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah

	fmt.Println("\n\n=========== Single pointer 3")
	// More ways for pointer
	var address3 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address4 *Address = &address3

	address4.City = "Jakarta"
	fmt.Println(address3)
	fmt.Println(address4)

	fmt.Println("\n\n=========== All pointer")
	address5 := &address3
	address6 := &address3
	address7 := &address3

	*address4 = Address{"Batak", "Jawa Barat", "Indonesia"}
	// *address5 = Address{"Bekasi", "Jawa Barat", "Indonesia"} // coba nyalakan comment ini dan lihat hasilnya
	fmt.Println(address3) // berubah
	fmt.Println(address4) // berubah
	fmt.Println(address5) // berubah
	fmt.Println(address6) // berubah
	fmt.Println(address7) // berubah

	fmt.Println("\n\n=========== (new) keyword for create a new pointer")
	address8 := new(Address) // recommended for create a new pointer
	fmt.Println(address8)    // empty pointer from struct Address

	address8.City = "Sukabumi"
	fmt.Println(address8)

	address9 := address8
	address9.City = "Palu"
	fmt.Println(address8)
	fmt.Println(address9)
}
