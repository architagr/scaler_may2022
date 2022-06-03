package closest_minmax

import (
	"math"
	maxminarray "scaler-may-22/week3/day8_assingment/max_min_array"
)

func FindClosestMinMax(A []int) int {
	length, ans, minIndex, maxIndex := 0, len(A), -1, -1
	max, min := maxminarray.MaxMinArray(A)

	if min == max {
		return 1
	}
	for i := len(A) - 1; i > -1; i-- {
		if A[i] == min {
			minIndex = i
		}
		if A[i] == max {
			maxIndex = i
		}
		if minIndex != -1 && maxIndex != -1 {
			length = int(math.Abs(float64(maxIndex-minIndex))) + 1
			if length < ans {
				ans = length
			}
		}

	}
	return ans
}
