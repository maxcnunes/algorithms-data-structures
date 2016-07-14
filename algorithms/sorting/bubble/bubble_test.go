package bubble

import "testing"

func TestSort(t *testing.T) {
	list := []int{8, 3, 6, 1, 5, 9, 7, 2, 4, 10}

	Sort(list)

	for i := 0; i < 10; i++ {
		if list[i] != i+1 {
			t.Error("Buble got wrong sort result", list)
		}
	}
}
