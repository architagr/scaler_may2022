package maximum_submatrix_sum_sorted_matrix

func FindMaxSubmatrixSum(A [][]int) int64 {
	n, m := len(A), len(A[0])
	arr := make([][]int64, n)
	row := make([]int64, m)

	if A[n-1][m-1] < 0 {
		return int64(A[n-1][m-1])
	}
	row[m-1] = int64(A[n-1][m-1])
	max := int64(A[n-1][m-1])
	for j := m - 2; j >= 0; j-- {
		row[j] = row[j+1] + int64(A[n-1][j])
		max = maxValue(max, row[j])
	}
	arr[n-1] = row
	for i := n - 2; i >= 0; i-- {
		row = make([]int64, m)
		row[m-1] = int64(A[i][m-1])
		for j := m - 2; j >= 0; j-- {
			row[j] = row[j+1] + int64(A[i][j])
		}
		arr[i] = row
	}
	for i := n - 2; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			arr[i][j] += arr[i+1][j]
			max = maxValue(max, arr[i][j])
		}
	}
	return max
}
func maxValue(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
