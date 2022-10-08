package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "hello"
	helloClone := strings.Clone(hello)
	fmt.Println("Clone :", helloClone)

	fmt.Println("Compare1 :", strings.Compare("a", "b")) // -1
	fmt.Println("Compare2 :", strings.Compare("a", "a")) // 0
	fmt.Println("Compare3 :", strings.Compare("b", "a")) // 1
	fmt.Println("Compare4 :", strings.Compare("a", "c")) // -1
	fmt.Println("Compare5 :", strings.Compare("c", "a")) // 1

	/*
		Kenapa Compare returnnya string?
		"Karena Compare function yang ada pada package strings ini adalah mengkomper secara leksikografis"

		Contoh:
			fmt.Println(strings.Compare("a", "b")) -1 atau tidak sama
			fmt.Println(strings.Compare("a", "a")) 0 atau tidak sama
			fmt.Println(strings.Compare("b", "a")) 1 atau tidak sama (kenapa gak negativ 1 kita tau urutan alphabet itu di mulai dari a)

		Contoh lain:
			a := "Hello"
			b := "Hello"
			The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	*/

	fmt.Println("Contains :", strings.Contains("Yadi Apriyadi", "Yadi"))             // true
	fmt.Println("Split1 :", strings.Split("Yadi Apriyadi", " "))                     // [Yadi Apriyadi] <= slice of string
	fmt.Println("Split2 :", strings.Split("Yadi Apriyadi", ""))                      // [Y a d i A p r i y a d i] <= slice of string
	fmt.Println("Split3 :", strings.Split("Yadi-Apriyadi", "-"))                     // [Yadi Apriyadi] <= slice of string
	fmt.Println("ToLower :", strings.ToLower("Yadi Apriyadi"))                       // yadi apriyadi
	fmt.Println("ToUpper :", strings.ToUpper("Yadi Apriyadi"))                       // YADI APRIYADI
	fmt.Println("ToTitle :", strings.ToTitle("yadi apriyadi"))                       // YADI APRIYADI
	fmt.Println("Trim :", strings.Trim("        yadi apriyadi    ", " "))            // yadi apriyadi <= remove space not a tab space!
	fmt.Println("ReplaceAll :", strings.ReplaceAll("yadi apriyadi", "yadi", "udin")) // udin apriudin
}
