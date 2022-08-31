package rotated_sorted_array_search

func RotatedSortedArray(A []int, B int) int {
	pivotIndex := findPivotIndex(A)
	l, r := 0, len(A)-1
	if A[0] < B {
		r = pivotIndex
	} else if A[0] > B {
		l = pivotIndex + 1
	} else {
		return 0
	}
	for l <= r {
		mid := (l + r) / 2
		if A[mid] < B {
			l = mid + 1
		} else if A[mid] > B {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func findPivotIndex(A []int) int {
	pivotIndex := 0
	l, r := 0, len(A)-1

	for l <= r {
		mid := (l + r) / 2
		if A[mid] > A[0] {
			l = mid + 1
			pivotIndex = mid
		} else {
			r = mid - 1
		}
	}
	return pivotIndex
}
