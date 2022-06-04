package rotate_matrix

func RotateMatrix(A *[][]int) {
	n := len(*A)
	for i := 0; i < ((n + 1) / 2); i++ {
		rotateBoundries(A, i, (n - 2*i))
	}
}

// for the a matrix of size 5 X 5 I have mentioned the index to be more clear
func rotateBoundries(A *[][]int, start, width int) {
	n := len(*A)
	t1, t2, t3, t4 := 0, 0, 0, 0

	for k := 0; k < width-1; k++ {
		// 0 0 | 0 1 | 0 2 | 0 3
		// 1 1 | 1 2
		t1 = (*A)[start][start+k]
		// 0 4 | 1 4 | 2 4 | 3 4
		// 1 3 | 2 3
		t2 = (*A)[start+k][n-start-1]

		// 4 4 | 4 3 | 4 2 | 4 1
		// 3 3 | 3 2
		t3 = (*A)[n-start-1][n-start-k-1]

		// 4 0 | 3 0 | 2 0 | 1 0
		// 3 1 | 2 1
		t4 = (*A)[n-start-k-1][start]

		(*A)[start][start+k] = t4
		(*A)[k+start][n-start-1] = t1
		(*A)[n-start-1][n-start-k-1] = t2
		(*A)[n-start-k-1][start] = t3
	}

}
