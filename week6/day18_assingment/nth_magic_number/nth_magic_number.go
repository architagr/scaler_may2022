package nth_magic_number

func FindNthMagicNumber(A int) int {
	fiveMultiple := 1
	ans := 0
	for i := 1; i <= 13; i++ {
		fiveMultiple *= 5
		test := 1 << (i - 1)
		if A&test == test {
			ans += fiveMultiple
		}
	}
	return ans
}
