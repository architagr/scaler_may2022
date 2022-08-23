package kth_symbol

func KthSymbol(A, B int) int {
	return calc(A-1, B-1)
}
func calc(A, B int) int {
	if A == 0 || B == 0 {
		return 0
	}
	x := calc(A-1, B/2)
	if B&1 != 1 { // if even
		return x
	}
	return 1 ^ x
}
