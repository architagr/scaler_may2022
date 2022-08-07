package reverse_bits

func ReverseBits(num uint32) uint32 {
	ans := uint32(0)

	for i := 0; i < 32; i++ {
		if num&(1<<i) != 0 {
			ans += 1 << (31 - i)
		}
	}
	return ans
}
