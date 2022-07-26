package max_sum_contiguous_subarray

import "scaler-may-22/common"

func MaxSubArray(A []int) int {
	sum, max := 0, common.MinInt()

	for i := 0; i < len(A); i++ {
		sum += A[i]

		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
