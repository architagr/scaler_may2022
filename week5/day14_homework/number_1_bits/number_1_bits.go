package number_1_bits

func findNumerofOnesBits(A int) int {
	count := 0
	for A > 0 {
		if A%2 == 1 {
			count++
		}
		A /= 2
	}
	return count
}
