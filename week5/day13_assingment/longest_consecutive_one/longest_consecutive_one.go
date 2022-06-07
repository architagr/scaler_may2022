package longest_consecutive_one

func GetLongestConsecutiveOne(A string) int {
	n := len(A)
	countOne := 0
	maxLen := 0
	for i := 0; i < n; i++ {
		if string(A[i]) == "1" {
			countOne++
		}
	}

	for i := 0; i < n; i++ {
		left := 0
		for j := i - 1; j >= 0; j-- {
			if string(A[j]) == "1" {
				left++
			} else {
				break
			}
		}
		right := 0
		for j := i + 1; j < n; j++ {
			if string(A[j]) == "1" {
				right++
			} else {
				break
			}
		}
		x := left + right
		if countOne == x {
			if maxLen < x {
				maxLen = x
			}
		} else {
			if maxLen < (x + 1) {
				maxLen = (x + 1)
			}
		}
	}
	return maxLen
}
