package valid_sudoku

func ValidSudoku(A []string) int {
	horizontal := make(map[int]map[byte]int)
	vertical := make(map[int]map[byte]int)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if A[i][j] != '.' {
				// check in horizontal line if the value is there
				if row, ok := horizontal[i]; ok {
					if _, ok1 := row[A[i][j]]; ok1 {
						return 0
					} else {
						row[A[i][j]] = j
						horizontal[i] = row
					}
				} else {
					horizontal[i] = map[byte]int{A[i][j]: j}
				}
				// check in vertical line if the value is there
				if cloumn, ok := vertical[j]; ok {
					if _, ok1 := cloumn[A[i][j]]; ok1 {
						return 0
					} else {
						cloumn[A[i][j]] = i
						vertical[j] = cloumn
					}
				} else {
					vertical[j] = map[byte]int{A[i][j]: i}
				}
				//check in the block
				startI := i - (i % 3)
				startJ := j - (j % 3)
				for ; startI < i; startI++ {
					row := horizontal[startI]
					if x, ok1 := row[A[i][j]]; ok1 {
						if x >= startJ && x <= startJ+2 {
							return 0
						}
					}
				}
			}
		}
	}

	return 1
}
