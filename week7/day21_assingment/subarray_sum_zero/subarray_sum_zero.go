package subarray_sum_zero

func FindSubArrayWithSum0(A []int) int {
	preFixSum := findPrefixSum(A)
	n := len(A)
	flag := false
	for i := 0; i < n; i++ {
		x := int64(0)
		if i > 0 {
			x = preFixSum[i-1]
		}
		for j := i; j < n; j++ {
			sum := preFixSum[j] - x

			if sum == 0 {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	if flag {
		return 1
	}
	return 0

}

func findPrefixSum(A []int) []int64 {
	sum := int64(0)
	n := len(A)
	result := make([]int64, 0)
	for i := 0; i < n; i++ {
		sum += int64(A[i])
		result = append(result, sum)
	}
	return result
}
