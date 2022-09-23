package longest_substring_without_repeat

func LengthOfLongestSubstring(A string) int {

	l, r, ans := 0, 1, 1
	hashmap := make(map[byte]int)
	hashmap[A[l]] = 0
	for ; r < len(A); r++ {
		if i, ok := hashmap[A[r]]; ok && i >= l {
			ans = maxValue(ans, (r - l))
			l = i + 1
		}
		hashmap[A[r]] = r
	}
	ans = maxValue(ans, (r - l))

	return ans
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
