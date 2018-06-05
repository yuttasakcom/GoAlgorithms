package addnumber

import "testing"

func TestAddNumber(t *testing.T) {
	s1 := add(1, 2)
	if s1 != 3 {
		t.Errorf("Expected 3, but got %v", s1)
	}

	s2 := add(3, 2)
	if s2 != 5 {
		t.Errorf("Expected 5, but got %v", s2)
	}

	s3 := add(13, 4)
	if s3 != 17 {
		t.Errorf("Expected 17, but got %v", s3)
	}

	s4 := add(11, 12)
	if s4 != 23 {
		t.Errorf("Expected 23, but got %v", s4)
	}

	s5 := add(5, 7)
	if s5 != 12 {
		t.Errorf("Expected 12, but got %v", s5)
	}
}
