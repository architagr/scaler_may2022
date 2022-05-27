package equilibrium_index_of_array

import (
	"log"
	"scaler-may-22/week2/day9_assingment/range_sum_query"
)

func FindEquilibrimIndex(A []int) int {
	prefixSumArray := range_sum_query.FindPrefixSum(A)
	log.Printf("prefix array %+v", prefixSumArray)
	var leftSum int64 = 0
	var rightSum int64 = 0

	for i := 0; i < len(A); i++ {
		leftSum, rightSum = 0, 0
		if i > 0 {
			leftSum = prefixSumArray[i-1]
		}

		if i < (len(A) - 1) {
			rightSum = prefixSumArray[len(A)-1] - prefixSumArray[i]
		}
		if rightSum == leftSum {
			return i
		}
	}
	return -1
}
