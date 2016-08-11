package set

import "testing"

func TestSet(t *testing.T) {
	s := NewSet()
	if s == nil {
		t.Error("Set has not been create correctly", s)
		return
	}

	if s.Size != 0 {
		t.Errorf("Expected to init with 0 but got %d", s.Size)
		return
	}

	s.Add("Bob")
	if !s.Contains("Bob") || s.Size != 1 {
		t.Errorf("Expected to add new item and increase the size to 1 but got %d", s.Size)
		return
	}

	s.Add("Jack")
	if !s.Contains("Jack") || s.Size != 2 {
		t.Errorf("Expected to add new item and increase the size to 2 but got %d", s.Size)
		return
	}

	s.Add("Bob")
	if !s.Contains("Bob") || s.Size != 2 {
		t.Errorf("Expected to not add an existing item and increase the size to 3 but got %d", s.Size)
		return
	}

	s.Remove("Bob")
	if s.Contains("Bob") || s.Size != 1 {
		t.Errorf("Expected to remove existing item and decrease the size to 1 but got %d", s.Size)
		return
	}
}
