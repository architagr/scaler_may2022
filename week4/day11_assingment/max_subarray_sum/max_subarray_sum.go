package max_subarray_sum

func FindMaxSubArraySumBelowValue(A, B int, C []int) int {
	max := 0
	sum := 0

	for i := 0; i < A; i++ {
		sum = 0

		for j := i; j < A; j++ {
			sum += C[j]

			if sum <= B && sum > max {
				max = sum
			}
		}
	}
	return max
}
