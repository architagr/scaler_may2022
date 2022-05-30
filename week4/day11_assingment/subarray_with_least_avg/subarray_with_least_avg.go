package subarray_with_least_avg

import (
	"scaler-may-22/common"
	"scaler-may-22/week2/day9_assingment/range_sum_query"
)

func SubArrayWithLeastAverage(A []int, B int) int {
	prefixSum := range_sum_query.FindPrefixSum(A)
	min := common.MaxInt()
	index := -1
	for i := 0; i <= len(A)-B; i++ {
		sum := findQuerySum(prefixSum, i, i+B)
		if min > sum {
			min = sum
			index = i
		}
	}
	return index
}

func findQuerySum(prefixSumArray []int64, start, end int) int {
	var ans int64 = 0
	left := start
	right := end - 1
	ans = prefixSumArray[right]
	if left > 0 {
		ans -= prefixSumArray[left-1]
	}
	return int(ans)
}
