package subset

import (
	"sort"
)

func checkSet(n, k int) int {
	return n & (1 << (k - 1))
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

	if len(a) > len(b) {
		return false
	}

	return true

}
func GetSubset(A []int) (arr [][]int) {
	sort.Ints(A)
	for i := 0; i < (1 << len(A)); i++ {
		a := []int{}
		for j := 0; j < len(A); j++ {

			if checkSet(i, j+1) != 0 {
				a = append(a, A[j])
			}
		}
		arr = append(arr, a)
	}

	sort.Slice(arr, func(i, j int) bool { return sortCustome(arr, i, j) })
	return
}
