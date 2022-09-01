package matrix_search

func SearchMatrix(A [][]int, B int) int {
	n, m := len(A), len(A[0])
	r, c := n-1, 0

	for r >= 0 && c < m {
		if A[r][c] > B {
			r--
		} else if A[r][c] < B {
			c++
		} else {
			return 1
		}
	}
	return 0
}
