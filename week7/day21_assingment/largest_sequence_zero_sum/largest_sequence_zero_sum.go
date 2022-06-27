package largest_sequence_zero_sum

func FindLsZero(A []int) []int {
	preFixSum := findPrefixSum(A)
	n := len(A)
	max := 0
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		x := int64(0)
		if i > 0 {
			x = preFixSum[i-1]
		}
		for j := i; j < n; j++ {
			sum := preFixSum[j] - x

			if sum == 0 && max < (j-i+1) {
				max = (j - i + 1)
				if j == n-1 {
					result = A[i:]
				} else {
					result = A[i : j+1]
				}
			}
		}
	}
	return result
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
