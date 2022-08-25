package maxmod

import "math"

func MaxMod(A []int) int {
	max, max2 := math.MinInt32, math.MinInt32

	for j := 0; j < len(A); j++ {
		if A[j] > max {
			max = A[j]
		}
	}
	for j := 0; j < len(A); j++ {
		if A[j] > max2 && A[j] < max {
			max2 = A[j]
		}
	}
	if max2 == math.MinInt32 {
		return 0
	}
	return max2 % max
}
