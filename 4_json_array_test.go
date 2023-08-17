package learn11json_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/YadBro/belajar-golang/tree/11.-golang-json/helper"
	"github.com/YadBro/belajar-golang/tree/11.-golang-json/model"
)

func TestJsonArrayEncode(t *testing.T) {
	customerBudi := model.Customer{
		Firstname: "Budi",
		Lastname:  "Baik",
		Age:       1,
		Hobbies:   []string{"baca qur'an", "sapu halaman depan", "berbagi uang"},
	}

	helper.LogJson(customerBudi)
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"firstName":"Budi","lastName":"Baik","age":1,"hobbies":["baca qur'an","sapu halaman depan","berbagi uang"]}`
	jsonBytes := []byte(jsonString)

	customerBudi := &model.Customer{}

	json.Unmarshal(jsonBytes, customerBudi)

	fmt.Println(customerBudi)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customerJoko := &model.Customer{
		Firstname: "Joko",
		Lastname:  "Harimawan",
		Age:       15,
		Hobbies:   []string{"Fishing", "Cloud Computing"},
		Addresses: []model.Address{
			{
				Street:     "Jalan ABC",
				Country:    "Indonesia",
				PostalCode: "405393",
			},
			{
				Street:     "Jalan ABC 2",
				Country:    "Indonesia",
				PostalCode: "405393",
			},
		},
	}

	helper.LogJson(customerJoko)
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"firstName":"Joko","lastName":"Harimawan","age":15,"hobbies":["Fishing","Cloud Computing"],"Addresses":[{"Street":"Jalan ABC","Country":"Indonesia","PostalCode":"405393"},{"Street":"Jalan ABC 2","Country":"Indonesia","PostalCode":"405393"}]}`
	jsonBytes := []byte(jsonString)

	customerJoko := &model.Customer{}

	json.Unmarshal(jsonBytes, customerJoko)

	fmt.Println(customerJoko)
	fmt.Println(customerJoko.Firstname)
	fmt.Println(customerJoko.Lastname)
	fmt.Println(customerJoko.Age)
	fmt.Println(customerJoko.Hobbies)
	fmt.Println(customerJoko.Addresses[0])
	fmt.Println(customerJoko.Addresses[0].Country)
	fmt.Println(customerJoko.Addresses[0].PostalCode)
	fmt.Println(customerJoko.Addresses[0].Street)
	fmt.Println(customerJoko.Addresses[1])
	fmt.Println(customerJoko.Addresses[1].Country)
	fmt.Println(customerJoko.Addresses[1].PostalCode)
	fmt.Println(customerJoko.Addresses[1].Street)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan ABC","Country":"Indonesia","PostalCode":"405393"},{"Street":"Jalan ABC 2","Country":"Indonesia","PostalCode":"405393"}]`
	jsonBytes := []byte(jsonString)

	customerJoko := &[]model.Address{}

	json.Unmarshal(jsonBytes, customerJoko)

	fmt.Println(customerJoko)
}
