package row_to_column_zero

func RowToColumnZero(A [][]int) [][]int {
	n := len(A)
	m := len(A[0])
	c := Copy(A)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				for jt := 0; jt < m; jt++ {
					c[i][jt] = 0
				}
				for it := 0; it < n; it++ {
					c[it][j] = 0
				}
			}
		}
	}
	return c
}

func Copy(A [][]int) [][]int {
	n := len(A)
	m := len(A[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i] = append(result[i], A[i][j])
		}
	}

	return result

}
