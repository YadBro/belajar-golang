package helper_session2

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// Access Modifier / Export
/*
	Jadi access modifier simplenya adalah seperti kita ingin mengexport sebuah function ataupun variable, dsb.
	Untuk melakukan access modifer atau export di golang cukup membuat sebuah function, variable, dsb. Dengan
	diawali oleh huruf besar
*/

const Application = "Belajar Golang"

// WARNING!!!
var version = 1      // ini saja sudah diberi warning (var version is unused (U1000))
const versionApi = 2 // ini saja sudah diberi warning (const version is unused (U1000))

// Jadi kalo mau buat package terus pengen ngexport itu harus diawali dengan huruf besar
const Author = "Yadi Apriyadi"
