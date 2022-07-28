package search_in_a_row_wise_and_column_wise_sorted_matrix

import "scaler-may-22/common"

func Search(A [][]int, B int) int {
	i, j := 0, len(A[0])-1
	MAX_INT := common.MaxInt()
	maxPos := MAX_INT
	for i < len(A) && j >= 0 {
		if A[i][j] == B {
			x := ((i + 1) * 1009) + (j + 1)
			if x < maxPos {
				maxPos = x
			}
			j--
		} else if A[i][j] > B {
			j--
		} else if A[i][j] < B {
			i++
		}
	}
	if maxPos == MAX_INT {
		return -1
	}
	return maxPos
}
