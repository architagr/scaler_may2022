package tower_of_hanoi

func TowerOfHanoi(A int) [][]int {
	result := make([][]int, 0)
	toh(A, 1, 2, 3, &result)
	return result
}
func toh(A, source, helper, destination int, result *[][]int) {
	if A == 0 {
		return
	}
	toh(A-1, source, destination, helper, result)
	*result = append(*result, []int{A, source, destination})
	toh(A-1, helper, source, destination, result)
}
