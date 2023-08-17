package helper

import (
	"encoding/json"
	"fmt"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data) // Merubah hasil golang ke json (encoding)

	PanicIfErr(err)
	fmt.Println(string(bytes))
}
