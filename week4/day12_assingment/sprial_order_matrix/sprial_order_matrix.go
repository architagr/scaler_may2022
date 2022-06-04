package sprial_order_matrix

func GetSprialOrderMatrix(A int) [][]int {
	result := createEmptyArray(A)
	startNu := 1
	for i := 0; i < ((A + 1) / 2); i++ {
		startNu = fillArray(&result, i, i, (A - 2*i), startNu)
	}
	return result
}
func fillArray(arr *[][]int, top, left, width, startNo int) int {
	i := top
	j := left

	for ; j < (left + width); j++ {
		(*arr)[i][j] = startNo
		startNo++
	}
	i++
	j--
	for ; i < (top + width); i++ {
		(*arr)[i][j] = startNo
		startNo++
	}
	i--
	j--
	for ; j >= left; j-- {
		(*arr)[i][j] = startNo
		startNo++
	}
	i--
	j++
	for ; i > top; i-- {
		(*arr)[i][j] = startNo
		startNo++
	}
	return startNo
}

func createEmptyArray(a int) [][]int {
	result := make([][]int, a)
	for i := 0; i < a; i++ {
		sub := make([]int, a)
		result[i] = sub
	}

	return result
}
