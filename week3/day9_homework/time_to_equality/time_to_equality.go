package time_to_equality

func TimeToEquality(A []int) int {
	max := A[0]
	count := 0
	for i, val := range A {
		if val > max {
			count += (val - max) * i
			max = val
		} else {
			count += max - val
		}
	}

	return count
}
