package sieveoferatosthenes

import "testing"

func TestFindPrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	actual := FindPrimes(30)

	if len(expected) != len(actual) {
		t.Errorf(
			"FindPrimes got wrong result for 30. Expected %v - Actual %v",
			expected,
			actual,
		)
		return
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf(
				"FindPrimes got wrong result for 30. Expected %v - Actual %v",
				expected,
				actual,
			)
		}
	}
}
