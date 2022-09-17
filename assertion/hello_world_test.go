package assertion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Yadi")

	// assert.Equal(t, "Hello Yadi", result, "Result must be 'Hello Yadi'") // success
	// assert.Equal(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // fail
	// assert.NotEqual(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // success
	assert.Nil(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // success

	fmt.Println("TestHelloWorld with done!") // di eksekusi walaupun fail
}

func TestHelloWorldAssertNil(t *testing.T) {
	yadi := Person{
		Name: "",
	}
	result := HelloWorldObject(yadi)

	// assert.Equal(t, "Hello Yadi", result, "Result must be 'Hello Yadi'") // success
	// assert.Equal(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // fail
	// assert.NotEqual(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // success

	assert.Nil(t, result, "Result must be 'nil'") // success
	if assert.NotNil(t, result) == true {
		assert.Equal(t, "Yadi", result, "Result must be 'Yadi'")
		fmt.Println("OK")
	}

	fmt.Println("HelloWorldObject with done!") // di eksekusi walaupun fail
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Yadi")

	// require.Equal(t, "Hello Yadi", result, "Result must be 'Hello Yadi'") // success
	// require.Equal(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // fail
	require.NotEqual(t, "Hi Yadi", result, "Result must be 'Hello Yadi'") // success

	fmt.Println("Tidak di eksekusi!") // tidak di eksekusi apabila fail
}
