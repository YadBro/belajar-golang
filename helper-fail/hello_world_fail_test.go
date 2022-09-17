package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldYadi(t *testing.T) {
	result := HelloWorldFail("Yadi")

	if result != "Hello Yadi" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorld Yadi done")
}

func TestHelloWorldApriyadi(t *testing.T) {
	result := HelloWorldFail("Apriyadi")

	if result != "Hello Apriyadi" {
		t.FailNow()
	}

	fmt.Println("TestHelloWorld Apriyadi done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorldFail("Yadi")

	if result != "Hello Yadi" {
		t.Error("Result must be 'Hello Yadi'")
	}

	fmt.Println("TestHelloWorldError Yadi done")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorldFail("Apriyadi")

	if result != "Hello Apriyadi" {
		t.Fatal("Result must be 'Hello Apriyadi'")
	}

	fmt.Println("TestHelloWorldError Apriyadi done")
}
