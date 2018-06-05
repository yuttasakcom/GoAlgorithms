package reverseint

import "testing"

func TestReverseInt(t *testing.T) {
	r1 := reverseInt(1234)
	if r1 != 4321 {
		t.Errorf("Expected 4321, but got %v", r1)
	}

	r2 := reverseInt(-1234)
	if r2 != -4321 {
		t.Errorf("Expected -4321, but got %v", r2)
	}

	r3 := reverseInt(4321)
	if r3 != 1234 {
		t.Errorf("Expected 1234, but got %v", r3)
	}
}
