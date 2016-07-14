package bubble

func Sort(items []int) {
	numItems := len(items)

	sorting := false
	for i := 0; i < numItems-1; i++ {
		for j := range items {
			if j == numItems-1 {
				continue
			}

			current := items[j]
			next := items[j+1]

			if current > next {
				sorting = true
				items[j] = next
				items[j+1] = current
			}
		}

		if !sorting {
			break
		}
	}
}
