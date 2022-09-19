package allocate_books

func Books(A []int, B int) int {
	n := len(A)
	if n < B {
		return -1
	}
	sum, max := 0, A[0]

	for i := 1; i < n; i++ {
		sum += A[i]
		if A[i] > max {
			max = A[i]
		}
	}
	if B == n {
		return max
	}
	l, r := max, 1000000000
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		if possible(A, mid, B) {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
func possible(arr []int, mid, B int) bool {
	sum := 0
	countofstuds := 1
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > mid {
			countofstuds++
			sum = arr[i]
		}
	}
	return countofstuds <= B
}
