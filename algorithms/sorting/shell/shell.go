package shell

// Sort - Shell has a better performance than the Bubble.
// Because it starts by bringing the items closer to where they actually belong.
// Example sorting with incremental sequence of 2 (Donald Shell).
// Given the disordered list:
// [8, 3, 6, 1, 5, 9, 7, 2, 4, 10] First interaction: int(10/2) = 5 sublists
// [8              9             ]
// [   3              7          ]
// [      6              2       ]
// [         1              4    ]
// [            5              10]
// [8, 3, 2, 1, 5, 9, 7, 6, 4, 10] Second interaction: int(5/2) = 2 sublists
// [8     2     5     7     4, 10]
// [   3     1     9     6     10]
// [2, 1, 4, 3, 5, 6, 7, 9, 8, 10] Third interaction: int(2/2) = 1 sublists
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] Result
func Sort(items []int, typeSort string) {
	switch typeSort {
	case "sedgewick":
		sortSedgewick(items)
	default:
		sortDonald(items)
	}
}

// Implementation based on Donald Shell approach.
// Seems it is the most known of implementation for this algorithm.
// References:
// - http://interactivepython.org/runestone/static/pythonds/SortSearch/TheShellSort.html
func sortDonald(items []int) {
	size := len(items)
	interval := size

	for interval > 0 {
		for idxSublist := 0; idxSublist < interval; idxSublist++ {
			for idxItem := idxSublist + interval; idxItem < size; idxItem = idxItem + interval {
				value := items[idxItem]
				pos := idxItem

				for pos >= interval && items[pos-interval] > value {
					items[pos] = items[pos-interval]
					pos = pos - interval
				}

				items[pos] = value
			}
		}

		interval = interval / 2
	}
}

// Implementation based on Robert Sedgewick approach.
// It has a better worst case performance than the previous one.
// References:
// - https://www.youtube.com/watch?v=9J14Pbw_XnY
// - http://www.tutorialspoint.com/data_structures_algorithms/shell_sort_algorithm.htm
func sortSedgewick(items []int) {
	size := len(items)
	interval := 1
	for interval < size/3 {
		interval = 3*interval + 1
	}

	for interval > 0 {
		for idxSublist := interval; idxSublist < size; idxSublist++ {
			value := items[idxSublist]
			idxItem := idxSublist

			for idxItem > interval-1 && items[idxItem-interval] >= value {
				items[idxItem] = items[idxItem-interval]
				idxItem = idxItem - interval
			}

			items[idxItem] = value
		}

		interval = (interval - 1) / 3
	}
}
