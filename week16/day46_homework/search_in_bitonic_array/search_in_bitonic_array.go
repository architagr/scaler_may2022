package search_in_bitonic_array

func BitonicArraySearch(A []int, B int) int {
	pivotIndex := findPivotIndex(A)
	index := searchAsc(A, 0, pivotIndex, B)
	if index == -1 {
		index = searchDesc(A, pivotIndex+1, len(A)-1, B)
	}
	return index
}
func searchAsc(A []int, l, r, b int) int {
	for l <= r {
		mid := (l + r) / 2
		if A[mid] == b {
			return mid
		} else if A[mid] < b {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func searchDesc(A []int, l, r, b int) int {
	for l <= r {
		mid := (l + r) / 2
		if A[mid] == b {
			return mid
		} else if A[mid] > b {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func findPivotIndex(A []int) int {
	l, r := 0, len(A)-1

	for l <= r {
		mid := (l + r) / 2
		if A[mid-1] < A[mid] && A[mid] < A[mid+1] {
			l = mid + 1
		} else if A[mid-1] > A[mid] && A[mid] > A[mid+1] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
