package square_root_of_integer

func Sqrt(A int) int {
	if A == 1 {
		return 1
	}
	l, r := 1, A
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == A {
			return mid
		}
		if mid*mid < A {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
