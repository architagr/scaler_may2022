package number_of_islands

/**
 * @input A : 2D integer array
 *
 * @Output Integer
 */
func NumberOfIsland(A [][]int) int {
	count := 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				dfs(A, i, j)
				count++
			}
		}
	}

	return count
}

func dfs(A [][]int, i, j int) {
	n, m := len(A), len(A[0])
	if i < 0 || j < 0 || j >= m || i >= n || A[i][j] == 0 {
		return
	}
	A[i][j] = 0
	dfs(A, i-1, j)
	dfs(A, i, j-1)
	dfs(A, i+1, j)
	dfs(A, i, j+1)

	dfs(A, i-1, j-1)
	dfs(A, i-1, j+1)
	dfs(A, i+1, j+1)
	dfs(A, i+1, j-1)
}
