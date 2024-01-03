package method

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {

	length := 4

	l := len(GenerateRandomString(length))
	if l != length {
		t.Errorf("GenerateRandomString(%v) return %v, expected %v", length, l, length)
	}
}
