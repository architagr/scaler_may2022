package flip

func CalcFlip(A string) []int {
	ans := make([]int, 2)
	max, sum, l, r := 0, 0, -1, -1
	for i := 0; i < len(A); i++ {
		if A[i] == '0' {
			sum++
			if l == -1 {
				l = i
				r = i
			}
		} else {
			sum--
		}
		if max < sum {
			max = sum
			r = i
			ans[0] = l + 1
			ans[1] = r + 1
		}

		if sum < 0 {
			sum = 0
			l = -1
		}
	}
	if l == -1 && ans[0] == 0 {
		return []int{}
	}
	return ans
}
