package subarray_with_given_sum

func SubarrayWithGivenSum(A []int, B int) []int {
	prefixSum := getPrefixSum(A)
	i, j := 0, 1

	for i <= j && j < len(A) {
		sum := prefixSum[j]
		if i > 0 {
			sum -= prefixSum[i-1]
		}

		if sum < B {
			j++
		} else if sum > B {
			i++
		} else {
			return A[i : j+1]
		}
	}
	return []int{-1}
}

func getPrefixSum(A []int) []int {
	prefixSum := make([]int, len(A))
	prefixSum[0] = A[0]
	for i := 1; i < len(A); i++ {
		prefixSum[i] = prefixSum[i-1] + A[i]
	}
	return prefixSum
}
