package number_of_1_bits

func Count1Bits(A int) int {
	count := 0

	for i := 0; i < 32; i++ {
		x := 1 << i
		if A&x != 0 {
			count++
		}
	}
	return count
}
