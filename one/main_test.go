package one

import (
	"testing"
)

func TestOne(t *testing.T) {
	result := One("../input/one/one-test.txt")
	if result != 7 {
		t.Log("Expected 7, got ", result)
		t.Fail()
	}
}
