package range_sum_query

func FindSumRange(A []int, B [][]int) []int64 {
	prefixSumArray := FindPrefixSum(A)
	result := make([]int64, len(B))
	var ans int64 = 0
	for i, query := range B {
		left := query[0] - 1
		right := query[1] - 1
		ans = prefixSumArray[right]

		if left != 0 {
			ans -= prefixSumArray[left-1]
		}
		result[i] = ans
	}
	return result
}

func FindPrefixSum(A []int) []int64 {
	result := make([]int64, len(A))
	var sum int64 = 0
	for i, val := range A {
		sum += int64(val)
		result[i] = sum
	}
	return result
}
