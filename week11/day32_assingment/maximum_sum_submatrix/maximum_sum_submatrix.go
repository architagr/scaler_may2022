package maximum_sum_submatrix

import "scaler-may-22/common"

func MaxSubMatrix(A [][]int) int {
	n := len(A)
	m := len(A[0])
	maxsum := common.MinInt()

	for k := 0; k < m; k++ {
		rowarr := make([]int, n)
		for i := k; i < m; i++ {
			for j := 0; j < n; j++ {
				rowarr[j] = rowarr[j] + A[j][i]
			}
			maxsum = maxValue(kadanes(rowarr), maxsum)
		}
	}
	return maxsum
}

func kadanes(rowarr []int) int {
	sum, maxsum := 0, common.MinInt()
	for i := 0; i < len(rowarr); i++ {
		sum = sum + rowarr[i]
		maxsum = maxValue(sum, maxsum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxsum
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
