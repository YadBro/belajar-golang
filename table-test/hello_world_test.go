package table_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TableHelloWorldTest []struct {
	Name, Request, Expected string
}

func TestTableTest(t *testing.T) {
	tests := TableHelloWorldTest{
		{
			Name:     "Yadi",
			Request:  "Yadi",
			Expected: "Hello Yadi",
		},
		{
			Name:     "Apriyadi",
			Request:  "Apriyadi",
			Expected: "Hello Apriyadi",
		},
		{
			Name:     "Eko",
			Request:  "Eko",
			Expected: "Hello Eko",
		},
		{
			Name:     "Khannedy",
			Request:  "Khannedy",
			Expected: "Hello Khannedy",
		},
		{
			Name:     "Kurniawan",
			Request:  "Kurniawan",
			Expected: "Hello Kurniawan",
		},
		{
			Name:     "Budi",
			Request:  "Budi",
			Expected: "Hello Budi",
		},
		{
			Name:     "Joko",
			Request:  "Joko",
			Expected: "Hello Joko",
		},
	}

	for _, value := range tests {
		t.Run(value.Name, func(t *testing.T) {
			result := HelloWorld(value.Request)
			require.Equal(t, value.Expected, result, "Result must be "+"'"+value.Expected+"'")
			fmt.Println("Run test " + value.Name + " done!")
		})
	}

}
