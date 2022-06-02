package sum_all_subarray

func SubarraySum(A []int) int64 {
	var total int64 = 0
	for i := range A {
		total += int64(A[i] * (len(A) - i) * (i + 1))
	}
	return total
}
