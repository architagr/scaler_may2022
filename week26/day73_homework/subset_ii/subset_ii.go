package subset_ii

/**
 * @input A : Integer array
 *
 * @Output 2D int array.
 */
import "sort"

var hashMap map[int]bool

func subsets(A []int) [][]int {
	sort.Ints(A)
	hashMap = make(map[int]bool)
	result := make([][]int, 0)
	arr := make([]int, 0)
	generate(arr, A, 0, &result)

	sort.Slice(result, func(i, j int) bool { return sortCustome(result, i, j) })
	return result
}
func sortCustome(arr [][]int, p int, q int) bool {
	a := arr[p]
	b := arr[q]

	for i := 0; i < len(a) && i < len(b); i++ {

		if a[i] < b[i] {
			return true
		}

		if a[i] > b[i] {
			return false
		}
	}

	return len(a) <= len(b)

}
func generate(arr []int, A []int, index int, result *[][]int) {
	if index == len(A) {
		num := createNum(arr)
		if _, ok := hashMap[num]; !ok {
			hashMap[num] = true
			*result = append(*result, append([]int{}, arr...))
		}
		return
	}
	generate(arr, A, index+1, result)
	generate(append(arr, A[index]), A, index+1, result)
}

func createNum(A []int) int {
	num := 0
	for _, val := range A {
		num = (num * 10) + val
	}
	return num
}
