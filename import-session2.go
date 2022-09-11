package main

import (
	"belajar-golang/helper_session2"
	"fmt"
)

func main() {
	helper_session2.SayHello("Yadi")

	// ACCESS MODIFER
	fmt.Println("Application:", helper_session2.Application)
	// fmt.Println(helper_session2.version) // error (panic)
	fmt.Println("Author:", helper_session2.Author)
}
