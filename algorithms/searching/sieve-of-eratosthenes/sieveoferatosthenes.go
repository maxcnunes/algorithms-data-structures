package sieveoferatosthenes

import (
	"math"
	"sort"
)

// FindPrimes will return the list of primes in the range between 2 and the maxNum.
// For large search is recommend to use the segmented approach instead of this basic one.
func FindPrimes(maxNum int) []int {
	nums := make(map[int]bool, maxNum-1)

	// fill list with the range of numbers
	for x := 2; x <= maxNum; x++ {
		nums[x] = true
	}

	// find out the prime numbers
	sqrt := int(math.Sqrt(float64(maxNum)))
	for i := 2; i <= sqrt; i++ {
		for j := 0; i*i+j*i <= maxNum; j++ {
			nums[i*i+j*i] = false
		}
	}

	// filter the prime numbers
	primes := []int{}
	for num, val := range nums {
		if val {
			primes = append(primes, num)
		}
	}

	// I could not find the reason.
	// But sometimes the result is not ordered.
	sort.Ints(primes)

	return primes
}
