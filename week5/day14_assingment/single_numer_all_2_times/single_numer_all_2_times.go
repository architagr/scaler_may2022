package single_numer_all_2_times

func FindSingleNumber(A []int) int {
	x := 0

	for _, val := range A {
		x ^= val
	}
	return x
}
