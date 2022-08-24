package kth_smallest_element

import "math"

func KthSmallestElement(A []int, B int) int {

	min := math.MaxInt32
	for i := 0; i < B; i++ {
		index := -1
		min = math.MaxInt32
		for j := i; j < len(A); j++ {
			if min > A[j] {
				min = A[j]
				index = j
			}
		}
		A[i], A[index] = A[index], A[i]
	}

	return min
}
