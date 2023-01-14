package uniqie_paths_3

func UniquePaths3(A [][]int) int {
	emptyCount, i1, j1 := 0, 0, 0
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				emptyCount++
			}
			if A[i][j] == 1 {
				i1 = i
				j1 = j
			}
		}
	}
	x := check(&A, 0, i1-1, j1, emptyCount)
	y := check(&A, 0, i1, j1-1, emptyCount)
	z := check(&A, 0, i1+1, j1, emptyCount)
	s := check(&A, 0, i1, j1+1, emptyCount)
	return x + y + z + s
}
func check(A *[][]int, count, i, j, emptyCount int) int {
	n, m := len((*A)), len((*A)[0])
	if i >= n || i < 0 || j >= m || j < 0 || (*A)[i][j] == -1 || (*A)[i][j] == 5 || (*A)[i][j] == 1 {
		return 0
	}
	if (*A)[i][j] == 2 {
		if emptyCount == count {
			return 1
		}
		return 0
	}
	(*A)[i][j] = 5
	x := check(A, count+1, i-1, j, emptyCount)
	y := check(A, count+1, i, j-1, emptyCount)
	z := check(A, count+1, i+1, j, emptyCount)
	s := check(A, count+1, i, j+1, emptyCount)
	(*A)[i][j] = 0
	return x + y + z + s
}
