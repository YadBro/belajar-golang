package learn11json_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/YadBro/belajar-golang/tree/11.-golang-json/helper"
	"github.com/YadBro/belajar-golang/tree/11.-golang-json/model"
)

func TestDecode(t *testing.T) {
	jsonString := `{"firstName":"Yadi","lastName":"Apriyadi","age":20}`
	jsonBytes := []byte(jsonString)

	customerYadi := &model.Customer{}

	err := json.Unmarshal(jsonBytes, customerYadi)
	helper.PanicIfErr(err)

	fmt.Println(customerYadi)
	fmt.Println(customerYadi.Firstname)
	fmt.Println(customerYadi.Lastname)
	fmt.Println(customerYadi.Age)
}
