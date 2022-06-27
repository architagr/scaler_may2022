package subarray_sum_zero

func FindSubArrayWithSum0(A []int) int {
	sum := int64(0)
	n := len(A)
	result := make(map[int64]int, 0)
	for i := 0; i < n; i++ {
		sum += int64(A[i])
		if _, ok := result[sum]; ok || sum == 0 {
			return 1
		} else {
			result[sum] = 1
		}
	}
	return 0
}
