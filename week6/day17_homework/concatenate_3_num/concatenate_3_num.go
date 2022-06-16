package concatenate_3_num

func solve(A int, B int, C int) int {
	abMin, s := min(A, B)
	abcMin, t := min(abMin, C)
	stMin, y := min(s, t)

	num := abcMin
	num = (num * 10000) + (stMin * 100) + y
	return num
}

func min(A, B int) (int, int) {
	if A < B {
		return A, B
	}
	return B, A
}
