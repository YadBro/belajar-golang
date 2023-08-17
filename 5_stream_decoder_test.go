package learn11json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/YadBro/belajar-golang/tree/11.-golang-json/helper"
	"github.com/YadBro/belajar-golang/tree/11.-golang-json/model"
)

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("data/product.json")
	helper.PanicIfErr(err)

	decoder := json.NewDecoder(reader)

	cilok := &model.Product{}
	errProduct := decoder.Decode(cilok)
	helper.PanicIfErr(errProduct)

	fmt.Println(cilok)
	fmt.Println(cilok.Name)
	fmt.Println(cilok.Category)
}
