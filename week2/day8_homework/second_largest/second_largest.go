package second_largest

func GetSecondHigest(A []int) int {
	higest, secondHigest := -1, -1

	for _, val := range A {
		if val > higest {
			secondHigest = higest
			higest = val
		} else if val > secondHigest {
			secondHigest = val
		}
	}
	return secondHigest
}
