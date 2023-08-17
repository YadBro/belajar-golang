package learn11json_test

import (
	"testing"

	"github.com/YadBro/belajar-golang/tree/11.-golang-json/helper"
)

func TestEncode(t *testing.T) {
	helper.LogJson("Yadi")
	helper.LogJson(1)
	helper.LogJson(true)
	helper.LogJson(false)
	helper.LogJson([]string{"Yadi", "Ganteng", "Sedunia"})
	helper.LogJson(map[string]interface{}{
		"Name":    "Yadi Ganteng Sedunia",
		"Title":   "Fullstack Developer",
		"Hobbies": []string{"Memasak", "Fishing", "Tidur", "Main Game"},
		"Social Media": map[string]interface{}{
			"Youtube": "https://www.youtube.com/channel/UCVb30GM1v3MEi3vWs48GAgg",
			"Github":  "https://github.com/YadBro",
		},
	})
}
