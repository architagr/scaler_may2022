package bit_compression

func compressBits(A []int) int {
	result := 0

	for _, val := range A {
		result ^= val
	}
	return result
}
