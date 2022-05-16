package squarerootofanumber

func SquareRoot(A int) int {
	var min, max, mid int = 1, A, 0

	for max >= min {
		mid = min + ((max - min) / 2)
		midSq := (mid * mid)
		if max-min <= 1 {
			if (max * max) == A {
				return max
			} else if (min * min) == A {
				return min
			}
		}
		if midSq == A {
			return mid
		} else if midSq > A {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}
