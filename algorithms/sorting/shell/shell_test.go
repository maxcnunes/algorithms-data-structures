package shell

import "testing"

func TestSortDonald(t *testing.T) {
	list := []int{8, 3, 6, 1, 5, 9, 7, 2, 4, 10}

	Sort(list, "donald")

	for i := 0; i < 10; i++ {
		if list[i] != i+1 {
			t.Error("Shell got wrong sort result", list)
			break
		}
	}
}

func TestSortSedgewick(t *testing.T) {
	list := []int{8, 3, 6, 1, 5, 9, 7, 2, 4, 10}

	Sort(list, "sedgewick")

	for i := 0; i < 10; i++ {
		if list[i] != i+1 {
			t.Error("Shell got wrong sort result", list)
			break
		}
	}
}
