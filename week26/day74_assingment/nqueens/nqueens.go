package nqueens

var result [][]string

func NQueens(n int) [][]string {
	result = make([][]string, 0)
	d1 := initDiagonal1Map(n)
	d2 := initDiagonal2Map(n)
	col := initColumnMap(n)
	solution := make([]string, n)
	for i := 0; i < n; i++ {
		solution[i] = initRow(n)
	}
	solve(0, n, solution, d1, d2, col)
	return result
}
func solve(row, n int, solution []string, d1, d2, col map[int]bool) {
	if row == n {
		x := make([]string, 0)
		x = append(x, solution...)
		result = append(result, x)
		return
	}
	for i := n - 1; i >= 0; i-- {
		if !col[i] || !d1[i-row] || !d2[row+i] {
			continue
		}
		rowConfig := []byte(solution[row])
		rowConfig[i] = 'Q'
		solution[row] = string(rowConfig)
		col[i] = false
		d1[i-row] = false
		d2[i+row] = false
		solve(row+1, n, solution, d1, d2, col)
		solution[row] = initRow(n)
		col[i] = true
		d1[i-row] = true
		d2[i+row] = true
	}
}
func initRow(n int) string {
	byteArr := make([]byte, n)
	for i := 0; i < n; i++ {
		byteArr[i] = '.'
	}
	return string(byteArr)
}
func initDiagonal1Map(n int) map[int]bool {
	diagonalMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		diagonalMap[i] = true
		diagonalMap[-i] = true
	}
	return diagonalMap
}
func initDiagonal2Map(n int) map[int]bool {
	diagonalMap := make(map[int]bool)
	for i := 0; i < (2*n)-1; i++ {
		diagonalMap[i] = true
	}
	return diagonalMap
}

func initColumnMap(n int) map[int]bool {
	columnMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		columnMap[i] = true
	}
	return columnMap
}
