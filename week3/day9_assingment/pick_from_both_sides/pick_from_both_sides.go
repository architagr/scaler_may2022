package pick_from_both_sides

import "scaler-may-22/week3/day9_assingment/range_sum_query"

func PickFromBothSides(A []int, B int) int {
	prefixSumArray := range_sum_query.FindPrefixSum(A)
	max := 0
	if B == len(A) {
		return int(prefixSumArray[len(A)-1])
	}
	for i := 0; i <= B; i++ {
		var leftSum int64 = 0
		var rightSum int64 = 0

		if i != 0 {
			leftSum = prefixSumArray[i-1]
		}

		if i != B {
			rightSum = prefixSumArray[len(A)-1] - prefixSumArray[len(A)-B+i-1]
		}
		sum := int(leftSum) + int(rightSum)
		if i == 0 {
			max = sum
		} else if max < sum {
			max = sum
		}
	}
	return max
}
