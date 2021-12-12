package two

import (
	"testing"
)

func TestOne(t *testing.T) {
	result := One("../input/two/one-test.txt")
	if result != 150 {
		t.Log("Expected 150, got ", result)
		t.Fail()
	}
}

// func TestTwo(t *testing.T) {
// 	result := Two("../input/two/two-test.txt")
// 	if result != 5 {
// 		t.Log("Expected 5, got ", result)
// 		t.Fail()
// 	}
// }
