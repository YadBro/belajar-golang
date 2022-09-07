package helper

import "fmt"

func GetGreeting2() string {
	return "HI"
}

var name2 = "udin terpesona"

func GetBro(name string) {
	fmt.Println("Hello Bro", name)
}

func main() {
	fmt.Println(name2)
}
