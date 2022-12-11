package max_product_subarray

import "math"

func maxProduct(A []int) int {
	max := math.MinInt
	for i := 0; i < len(A); i++ {
		prod := 1
		for j := i; j < len(A); j++ {
			prod *= A[j]
			max = maxValue(max, prod)
			if prod == 0 {
				break
			}
		}
	}
	return max
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
