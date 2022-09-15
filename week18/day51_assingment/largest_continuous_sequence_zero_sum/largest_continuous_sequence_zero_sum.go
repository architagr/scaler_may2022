package largest_continuous_sequence_zero_sum

func Lszero(A []int) []int {
	start, last, length := 0, 0, 0

	hashMap := make(map[int]int)
	hashMap[0] = -1
	cf := 0
	for i := 0; i < len(A); i++ {
		cf += A[i]
		if index, ok := hashMap[cf]; ok && i-index > length {
			start = index
			last = i
			length = last - start
		} else if !ok {
			hashMap[cf] = i
		}
	}

	return A[start+1 : last+1]
}
