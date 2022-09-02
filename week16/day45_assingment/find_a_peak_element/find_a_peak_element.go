package find_a_peak_element

func FindPeakElement(A []int) int {
	n := len(A)
	// check boundry elements

	if n == 1 {
		return A[0]
	}
	if A[n-1] > A[n-2] {
		return A[n-1]
	} else if A[0] > A[1] {
		return A[0]
	}

	l, r := 1, n-2
	for l <= r {
		mid := (l + r) / 2

		if A[mid] >= A[mid-1] && A[mid] >= A[mid+1] {
			return A[mid]
		}
		if A[mid] < A[mid-1] && A[mid] > A[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
