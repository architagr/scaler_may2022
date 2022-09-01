package sorted_insert_position

func SearchInsert(A []int, B int) int {
	l, r := 0, len(A)-1
	for l <= r {
		mid := (l + r) / 2
		if B < A[mid] {
			r = mid - 1
		} else if B > A[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}
