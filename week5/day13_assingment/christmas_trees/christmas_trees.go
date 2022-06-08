package christmas_trees

import "scaler-may-22/common"

func FindMinCostCristmasTrees(A, B []int) int {
	minCost := common.MaxInt()
	n := len(A)
	for i := 1; i < n; i++ {
		left := leftCost(i-1, A, B)
		right := rightCost(i+1, A, B)
		if left > 0 && right > 0 && (left+right-B[i]) < minCost {
			minCost = left + right - B[i]
		}

	}
	if minCost == common.MaxInt() {
		minCost = -1
	}
	return minCost
}

func leftCost(start int, A, B []int) int {
	min := common.MaxInt()
	for i := start; i >= 0; i-- {
		if A[i] < A[start+1] && B[start+1]+B[i] < min {
			min = B[start+1] + B[i]
		}
	}
	if min == common.MaxInt() {
		min = -1
	}
	return min
}
func rightCost(start int, A, B []int) int {
	min := common.MaxInt()
	for i := start; i < len(A); i++ {

		if A[i] > A[start-1] && B[start-1]+B[i] < min {
			min = B[start-1] + B[i]
		}
	}
	if min == common.MaxInt() {
		min = -1
	}
	return min
}
