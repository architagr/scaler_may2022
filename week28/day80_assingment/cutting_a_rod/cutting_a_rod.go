package cuttingarod

import "math"

func CuttingARod(A []int) int {
	n := len(A)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return rodHelper(A, n, dp)
}
func rodHelper(costs []int, length int, dp []int) int {
	if length <= 0 {
		return 0
	}
	if dp[length] != -1 {
		return dp[length]
	}
	maxCost := math.MinInt
	for i := 1; i <= len(costs); i++ {
		if (length - i) >= 0 {
			subProblemSolution := costs[i-1] + rodHelper(costs, length-i, dp)
			maxCost = maxValue(maxCost, subProblemSolution)
		}
	}
	dp[length] = maxCost
	return dp[length]
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
