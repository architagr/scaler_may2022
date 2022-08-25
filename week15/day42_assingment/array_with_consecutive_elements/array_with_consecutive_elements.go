package array_with_consecutive_elements

import "math"

func ArrayWithConsecutiveElements(A []int) int {

	min := math.MaxInt32

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
	}

	for i := 0; i < len(A); {
		if A[i] < min+len(A) && A[i] != min+i && A[i] != A[A[i]-min] {
			A[i], A[A[i]-min] = A[A[i]-min], A[i]

		} else {
			i++
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i] != min+i {
			return 0
		}
	}
	return 1
}
