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

func TestTwo(t *testing.T) {
	result := Two("../input/one/two-test.txt")
	if result != 5 {
		t.Log("Expected 5, got ", result)
		t.Fail()
	}
}
