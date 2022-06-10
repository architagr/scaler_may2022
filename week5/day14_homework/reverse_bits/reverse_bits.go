package reverse_bits

import "math"

func reverseBits(num uint32) uint32 {
	var x uint32 = 0
	max := uint32(math.Pow(2, 31))
	for num > 0 {
		r := num % 2
		x += r * max
		num /= 2
		max /= 2
	}
	return x
}
