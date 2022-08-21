package find_factorial

func Fact(A int) int {
	if A == 0 {
		return 1
	}
	return Fact(A-1) * A
}
