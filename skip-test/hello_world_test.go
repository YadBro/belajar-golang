package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

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
