package unset_bits

func UnsetBits(A int64, B int) int64 {

	x := MaxInt()

	x = x << B

	return (A & x)
}

func MaxInt() int64 {
	const MaxUint = ^uint64(0)
	const MaxInt = int64(MaxUint >> 1)
	return MaxInt
}
