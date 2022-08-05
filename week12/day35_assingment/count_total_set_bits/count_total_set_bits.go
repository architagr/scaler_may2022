package count_total_set_bits

func CountSetBits(A int) int {
	MOD := 1000000007
	if A == 0 {
		return 0
	}
	p := 0
	for ; (1 << p) <= A; p++ {
	}
	p--
	x := (1 << p)
	return (((((p * x / 2) % MOD) + ((A+1)%MOD - x%MOD)) % MOD) + (CountSetBits(A-x) % MOD)) % MOD
}
