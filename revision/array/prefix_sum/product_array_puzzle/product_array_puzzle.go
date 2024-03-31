package productarraypuzzle

func solve(A []int) []int {
	left, right, result := getLeftProductArray(A), getRightProductArray(A), make([]int, len(A))
	for i := 0; i < len(A); i++ {
		leftProduct, rightProduct := 1, 1
		if i > 0 {
			leftProduct = left[i-1]
		}
		if i+1 < len(A) {
			rightProduct = right[i+1]
		}
		result[i] = leftProduct * rightProduct
	}
	return result
}

func getLeftProductArray(input []int) []int {
	result := make([]int, len(input))
	result[0] = input[0]
	for i := 1; i < len(input); i++ {
		result[i] = result[i-1] * input[i]

	}
	return result
}

func getRightProductArray(input []int) []int {
	result, n := make([]int, len(input)), len(input)
	result[n-1] = input[n-1]
	for i := n - 2; i >= 0; i-- {
		result[i] = result[i+1] * input[i]

	}
	return result
}
