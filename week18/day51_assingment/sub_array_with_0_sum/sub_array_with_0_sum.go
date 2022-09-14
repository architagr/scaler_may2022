package sub_array_with_0_sum

func SubArray0Sum(A []int) int {
	prefixSum := make([]int, len(A))
	prefixSum[0] = A[0]
	for i := 1; i < len(A); i++ {
		prefixSum[i] = prefixSum[i-1] + A[i]
	}
	if prefixSum[len(prefixSum)-1] == 0 {
		return 1
	}
	hashmap := make(map[int]bool)

	for i := 0; i < len(prefixSum); i++ {
		if _, ok := hashmap[prefixSum[i]]; ok || prefixSum[i] == 0 {
			return 1
		}
		hashmap[prefixSum[i]] = true
	}

	return 0
}
