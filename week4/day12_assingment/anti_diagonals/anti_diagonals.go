package anti_diagonals

func GetAntiDiagonal(A [][]int) [][]int {
	n := len(A)
	result := make([][]int, 0)

	for j := 0; j < n-1; j++ {
		sub := make([]int, n)
		for i, ij, k := 0, j, 0; i < n-1 && ij >= 0; i, ij, k = i+1, ij-1, k+1 {
			sub[k] = A[i][ij]
		}
		result = append(result, sub)
	}

	for i := 0; i < n; i++ {
		sub := make([]int, n)
		for j, ji, k := n-1, i, 0; j >= 0 && ji < n; j, ji, k = j-1, ji+1, k+1 {
			sub[k] = A[ji][j]
		}
		result = append(result, sub)
	}

	return result
}
