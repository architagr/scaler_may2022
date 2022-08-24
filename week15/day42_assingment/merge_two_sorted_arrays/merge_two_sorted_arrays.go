package merge_two_sorted_arrays

func Merge2SortedArray(A, B []int) []int {
	i, j, k := 0, 0, 0
	n, m := len(A), len(B)
	result := make([]int, n+m)
	for ; i < n && j < m && k < (n+m); k++ {
		if A[i] < B[j] {
			result[k] = A[i]
			i++
		} else {
			result[k] = B[j]
			j++
		}
	}
	if i == n {
		for ; k < (n + m); k, j = k+1, j+1 {
			result[k] = B[j]
		}
	}
	if j == m {
		for ; k < (n + m); k, i = k+1, i+1 {
			result[k] = A[i]
		}
	}
	return result
}
