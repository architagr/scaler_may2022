package little_ponny_and_2Subsequence

func FindLexicographicalMinString(A string) string {
	var a, b byte = A[0], A[0]
	var aIndex, bIndex int = 0, 0

	for i := aIndex + 1; i < len(A)-1; i++ {
		if A[i] < a {
			a = A[i]
			aIndex = i
		}
	}

	b = A[aIndex+1]
	bIndex = aIndex + 1

	for i := bIndex + 1; i < len(A); i++ {
		if A[i] < b {
			b = A[i]
			bIndex = i
		}
	}
	return string([]byte{a, b})

}
