package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("y([a-z])i")

	fmt.Println(regex.MatchString("yai"))
	fmt.Println(regex.MatchString("yBi"))
	fmt.Println(regex.MatchString("yci"))
	fmt.Println(regex.MatchString("yei"))
	fmt.Println(regex.MatchString("yadi"))

	search := regex.FindAllString("yadi yoi yai yeo yci yoka", -1)
	fmt.Println(search)
}
