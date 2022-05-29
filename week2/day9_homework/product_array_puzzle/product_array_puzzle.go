package product_array_puzzle

func productArray(A []int) int {
	product := 1
	for _, val := range A {
		product *= val
	}
	return product
}

func ProductArrayPuzzle(A []int) []int {
	product := productArray(A)
	for i, val := range A {
		A[i] = product / val
	}
	return A
}
