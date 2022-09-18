package sort_array_in_given_order

import "sort"

func SortArray(A, B []int) []int {
	arr := make([]int, len(A))
	c := 0
	hashMap := make(map[int]int)
	sort.Ints(A)
	for _, val := range A {
		hashMap[val]++
	}
	for i := 0; i < len(B); i++ {
		if count, ok := hashMap[B[i]]; ok {
			for j := 0; j < count; j++ {
				arr[c] = B[i]
				c++
			}
			delete(hashMap, B[i])
		}
	}
	for i := 0; i < len(A); i++ {
		if _, ok := hashMap[A[i]]; ok {
			arr[c] = A[i]
			c++
		}
	}
	return arr
}
