package matrix_chain_multiplication

import "math"

var dp [][]int

func MatrixChainMultiplication(A []int) int {
	n := len(A)
	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	return checkValue(A, 1, len(A)-1)
}
func checkValue(A []int, i, j int) int {
	if i == j {
		dp[i][j] = 0

	} else if dp[i][j] == -1 {
		ans := math.MaxInt

		for k := i; k < j; k++ {
			x := checkValue(A, i, k)
			y := checkValue(A, k+1, j)
			z := A[i-1] * A[k] * A[j]
			ans = minValue(ans, x+y+z)
		}
		dp[i][j] = ans
	}
	return dp[i][j]
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
