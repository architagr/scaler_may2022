package row_with_maximum_number_of_ones

func MaxOnesRow(A [][]int) int {
	i, j := 0, len(A[0])-1
	maxJ, ans := j+1, -1
	for i < len(A) && j >= 0 {
		if A[i][j] == 1 {
			if j < maxJ {
				maxJ = j
				ans = i
			}
			j--
		} else {
			i++
		}
	}
	return ans
}
