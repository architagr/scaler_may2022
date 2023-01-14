package all_unique_permutations

var result [][]int
var hashMap map[int]bool

func permute(A []int) [][]int {
	result = make([][]int, 0)
	hashMap = make(map[int]bool)
	generate(A, 0)
	return result
}
func generate(A []int, start int) {
	if start == len(A) {
		num := createNum(A)
		if _, ok := hashMap[num]; !ok {
			hashMap[num] = true
			result = append(result, append(make([]int, 0), A...))
		}
		return
	}
	for i := start; i < len(A); i++ {
		A[i], A[start] = A[start], A[i]
		generate(A, start+1)
		A[i], A[start] = A[start], A[i]
	}
}
func createNum(A []int) int {
	num := 0
	for _, val := range A {
		num = (num * 10) + val
	}
	return num
}
