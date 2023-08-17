package learn11json_test

import (
	"testing"

	"github.com/YadBro/belajar-golang/tree/11.-golang-json/helper"
	"github.com/YadBro/belajar-golang/tree/11.-golang-json/model"
)

func TestJsonObject(t *testing.T) {
	customerYadi := model.Customer{
		Firstname: "Yadi",
		Lastname:  "Apriyadi",
		Age:       20,
	}

	helper.LogJson(customerYadi)
}
