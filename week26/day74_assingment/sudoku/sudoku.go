package sudoku

type Set map[int]bool
type ValueMap map[int]Set

func SolveSudoku(board *[][]byte) {
	row, col, cube := initMaps(board)
	findValue(board, 0, 0, row, col, cube)
}

func findValue(board *[][]byte, rowNo, colNo int, rowMap, colMap, cubeMap ValueMap) bool {
	if rowNo == 9 {
		return true
	}
	if colNo == 9 {
		return findValue(board, rowNo+1, 0, rowMap, colMap, cubeMap)
	}
	if (*board)[rowNo][colNo] == '.' {
		for i := 1; i <= 9; i++ {
			if !rowMap[rowNo][i] || !colMap[colNo][i] || !cubeMap[getCubeNo(rowNo, colNo)][i] {
				continue
			}
			rowMap[rowNo][i] = false
			colMap[colNo][i] = false
			cubeMap[getCubeNo(rowNo, colNo)][i] = false
			(*board)[rowNo][colNo] = byte(i) + '0'

			next := findValue(board, rowNo, colNo+1, rowMap, colMap, cubeMap)
			if !next {
				rowMap[rowNo][i] = true
				colMap[colNo][i] = true
				cubeMap[getCubeNo(rowNo, colNo)][i] = true
				(*board)[rowNo][colNo] = '.'
			}
		}
		if !validateRow(rowMap, rowNo) {
			return false
		}
	} else {
		return findValue(board, rowNo, colNo+1, rowMap, colMap, cubeMap)

	}
	return true
}
func validateRow(row ValueMap, rowNo int) bool {
	for _, val := range row[rowNo] {
		if val {
			return false
		}
	}
	return true
}
func initMaps(board *[][]byte) (rowMap ValueMap, colMap ValueMap, cubeMap ValueMap) {
	rowMap = make(ValueMap)
	colMap = make(ValueMap)
	cubeMap = make(ValueMap)

	for i := 0; i < 9; i++ {
		rowMap[i] = Set{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}
		colMap[i] = Set{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}
		cubeMap[i] = Set{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := (*board)[i][j]
			if val != '.' {
				n := int(val - '0')
				cubeNo := getCubeNo(i, j)
				rowMap[i][n] = false
				colMap[j][n] = false
				cubeMap[cubeNo][n] = false
			}
		}
	}
	return
}

func getCubeNo(rowNo, colNo int) int {
	return (3 * (rowNo / 3)) + (colNo / 3)
}
