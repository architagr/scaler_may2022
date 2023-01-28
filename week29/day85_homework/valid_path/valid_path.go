package valid_path

import "math"

func solve(x, y, N, R int, X, Y []int) string {
	mat := getMatrix(x, y, N, R, X, Y)

	if mat[0][0] == 0 || mat[x][y] == 0 {
		return "NO"
	}
	vis := make([][]bool, x+1)
	for i := 0; i <= x; i++ {
		vis[i] = make([]bool, y+1)
	}

	dfs(mat, 0, 0, vis)
	if !vis[x][y] {
		return "NO"
	}
	return "YES"
}
func dfs(mat [][]int, i, j int, vis [][]bool) {
	if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[0]) || mat[i][j] == 0 || vis[i][j] {
		return
	}
	vis[i][j] = true
	dfs(mat, i+1, j, vis)
	dfs(mat, i-1, j, vis)
	dfs(mat, i, j+1, vis)
	dfs(mat, i, j-1, vis)
	dfs(mat, i-1, j+1, vis)
	dfs(mat, i-1, j-1, vis)
	dfs(mat, i+1, j+1, vis)
	dfs(mat, i+1, j-1, vis)

}
func getMatrix(x, y, N, R int, X, Y []int) [][]int {
	mat := make([][]int, x+1)
	// create a initial matrix
	for i := 0; i <= x; i++ {
		mat[i] = make([]int, y+1)
	}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {

			//check all the circle points
			res := 1 //initially marked as 1
			for k := 0; k < N; k++ {

				//we are checking if the current cell is lying under the circle
				//boundary or not
				distance := int(math.Pow(float64(i-X[k]), float64(2)) + math.Pow(float64(j-Y[k]), float64(2)))
				if distance <= R*R {
					res = 0 //if distance less than the radius then set the cell to 0
					break
				}
			}
			mat[i][j] = res
		}
	}
	return mat
}
