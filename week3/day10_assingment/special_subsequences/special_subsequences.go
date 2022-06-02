package special_subsequences

func FindSpecialSubsequences(A string) int {
	countA := 0
	ans := 0

	for _, val := range A {
		if string(val) == "A" {
			countA++
		} else if string(val) == "G" {
			ans += countA
		}
	}
	return ans
}
