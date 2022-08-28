package main

import (
	"fmt"
)

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(1000 * time.Millisecond)
	// }

	// for i := 10; i >= 1; i-- {
	// 	fmt.Println(i)
	// 	time.Sleep(1000 * time.Millisecond)
	// }

	// fmt.Println("Selesai")

	// for i := 1; i <= 30; i++ {
	// 	if i%2 == 1 {
	// 		fmt.Println("Odd", i)
	// 	} else if i%3 == 1 && i%2 == 0 {
	// 		fmt.Println("Even", i)
	// 	}
	// }

	// apa:
	// 	for i := 0; i <= 5; i++ {
	// 		switch {
	// 		case i == 4:
	// 			fmt.Println("ganjil")
	// 		case i == 1:
	// 			fmt.Println("Ganjil")
	// 			break apa
	// 		default:
	// 			fmt.Println("Gak tau")
	// 		}
	// 	}

	// 	sistemOperasi := runtime.GOOS

	// 	fmt.Println(sistemOperasi)

	// data := Past(0, 1, 30)
	// fmt.Println(data)

	// for i := 0; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd", i)
	// 	if i == 15 {
	// 		break
	// 	}
	// }

	noBack := [3]string{"bambang", "asep", "udin"}
	students := [...]string{"Bambang", "Udin", "Rahmat"}
	fmt.Println(students)

	// for i := range noBack {
	// 	fmt.Println(noBack[i])
	// }

	fmt.Println("noBack", noBack)
	slices := noBack[:1]
	fmt.Println("slices", slices)
	fmt.Println("noBack", noBack)
	slices[0] = "terjindan"
	fmt.Println("slices", slices) // terjindan
	fmt.Println("noBack", noBack) //[terjindan asep udin]
	slices2 := noBack[1:]         // [asep udin]
	fmt.Println("slices2", slices2)
	slices2[1] = "terjindan jindan"
	fmt.Println("slices2", slices2)               // [asep terjindan jindan]
	slices2 = append(slices2, "watasi konodioda") //[asep terjindan jindan watasi konodioda]
	slices[0] = "pengacara pengacara"
	fmt.Println("slices", slices)
	fmt.Println("noBack", noBack)
	fmt.Println("slices2", slices2)

	// slice3 := make([]string, len(slices2))

	// copy(slice3, slices2)

	// fmt.Println(slice3)
	// slices[0] = "terpaku"
	// fmt.Println(slice3)

	// mahasiswa := map[string]map[string]string{
	// 	"001": {
	// 		"name":    "Udin",
	// 		"address": "Jl. Kenangan Mantan",
	// 		"gender":  "Male",
	// 	},
	// 	"002": {
	// 		"name":    "Yadi",
	// 		"address": "Jl. Kenangan Mantan",
	// 		"gender":  "Male",
	// 	},
	// 	"003": {
	// 		"name":    "Jamilah",
	// 		"address": "Jl. Kenangan Mantan",
	// 		"gender":  "Female",
	// 	},
	// }

	// for key, value := range mahasiswa {
	// 	delete(mahasiswa, key)
	// 	fmt.Println(key)
	// 	fmt.Println(value["name"])
	// }
	// fmt.Println(mahasiswa)

	// const name = "yadi"
	// fmt.Println(name)

	// var (
	// 	nilai32 int32
	// 	nilai64 int64
	// 	nilai8  int8
	// 	yadi    string
	// )
	// nilai32 = 127
	// nilai64 = int64(nilai32)
	// nilai8 = int8(nilai32)
	// yadi = "yadi"

	// char := yadi[0]

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)
	// fmt.Println(yadi)
	// fmt.Println(string(char))

	// var data4 string = "hello"
	// char1 := data4[0]
	// fmt.Println(string(char1))

	// var StringToNumber = strconv.Atoi
	// fmt.Println(StringToNumber("-7"))

	var StringToNumber = CountBy(2, 5)
	fmt.Println(StringToNumber)

}

func Past(h, m, s int) int {
	return (h*60*60 + m*60 + s) * 1000
}

func CountBy(x, n int) (data []int) {
	menujux := x * n // 250
	inc := x
	for x := x; inc <= menujux; inc += x {
		data = append(data, inc)
	}
	return data
}

func Maps(x []int) (result []int) {
	for _, value := range x {
		result = append(result, value*2)
	}
	return result
}

func multipleOfIndex(ints []int) (result []int) {
	for key, value := range ints {
		if key > 1 && key%value == 0 {
			result = append(result, ints[key])
		}
	}
	return result
}

func SmallestIntegerFinder(numbers []int) (min int) {
	min = numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
	}
	return min
}
