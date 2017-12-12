package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	s1 := reverseString("abcdef")
	if s1 != "fedcba" {
		t.Errorf("Expected fedcba, but got %v", s1)
	}

	s2 := reverseString("abc_")
	if s2 != "_cba" {
		t.Errorf("Expected _cba, but got %v", s2)
	}

	s3 := reverseString("_a_b_c_")
	if s3 != "_c_b_a_" {
		t.Errorf("Expected _c_b_a_, but got %v", s3)
	}
}
