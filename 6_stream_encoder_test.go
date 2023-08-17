package learn11json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/YadBro/belajar-golang/tree/11.-golang-json/helper"
	"github.com/YadBro/belajar-golang/tree/11.-golang-json/model"
)

func TestStreamEncoder(t *testing.T) {
	reader, err := os.Create("data/productHaus.json")
	helper.PanicIfErr(err)

	decoder := json.NewEncoder(reader)

	cilok := &model.Product{
		Name:     "Haus Jeruk",
		Category: "Drink",
	}
	errProduct := decoder.Encode(cilok)
	helper.PanicIfErr(errProduct)

	fmt.Println(cilok)
	fmt.Println(cilok.Name)
	fmt.Println(cilok.Category)
}
