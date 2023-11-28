package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		//panic("Result is not 'Hello Eko'")
		//t.Fail()
		t.Error("Result must be 'Hello Eko'")
	}
	fmt.Println("TestHelloWorldEko Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		//error
		//panic("Result is not 'Hello Khannedy'")
		//t.FailNow()
		t.Fatal("Result must be 'Hello Khannedy'")
	}
	fmt.Println("TestHelloWorldKhannedy Done")
}
