package main_test_before_and_after

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// BEFORE
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// AFTER
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows OS")
	}

	result := HelloWorld("Yadi")
	require.Equal(t, "Hello Yadi", result, "Result must be 'Hello Yadi'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yadi")

	if result != "Hello Yadi" {
		// error
		panic("Result is not Hello Yadi")
	}
}
