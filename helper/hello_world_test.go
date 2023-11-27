package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		panic("Result is not 'Hello Eko'")
	}
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		//error
		panic("Result is not 'Hello Khannedy'")
	}
}
