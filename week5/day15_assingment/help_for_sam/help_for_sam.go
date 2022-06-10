package help_for_sam

func FindHelpTakenFromSame(A int) int {
	count := 0
	for A > 0 {
		if A&1 == 1 {
			count++
		}
		A = A >> 1
	}

	return count
}
