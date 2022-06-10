package interesting_array

func FindIntrestingArray(A []int) string {
	count := 0

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 1 {
			count++
		}
	}
	if count%2 == 1 {
		return "No"
	}
	return "Yes"
}
