package maxminarray

func MaxMinArray(A []int) (int, int) {
	max := A[0]
	min := A[0]

	for _, val := range A {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	return max, min
}
