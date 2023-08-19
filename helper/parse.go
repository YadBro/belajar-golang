package helper

import "strconv"

func ParseToInt(str string) int {
	integer, err := strconv.Atoi(str)
	PanicIfError(err)

	return integer
}
