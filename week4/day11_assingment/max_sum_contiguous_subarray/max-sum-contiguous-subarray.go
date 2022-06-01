package max_sum_contiguous_subarray

import "scaler-may-22/common"

func GetMaxSumContiguousSubArray(A []int) int {
	intMin := common.MinInt()
	max := intMin
	n := len(A)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += A[j]

			if sum > max {
				max = sum
			}
		}
	}
	if max == intMin {
		max = 0
	}
	return max
}
