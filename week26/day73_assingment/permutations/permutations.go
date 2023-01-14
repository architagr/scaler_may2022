package permutations

/**
 * @input A : Integer array
 *
 * @Output 2D int array.
 */
func permute(A []int) [][]int {
	result := make([][]int, 0)
	generate(A, 0, &result)

	return result
}

func generate(A []int, index int, result *[][]int) {
	if index == len(A)-1 {
		arr := append(make([]int, 0, len(A)), A...)
		*result = append(*result, arr)
		return
	}
	for i := index; i < len(A); i++ {
		A[index], A[i] = A[i], A[index]
		generate(A, index+1, result)
		A[index], A[i] = A[i], A[index]
	}
	return
}
