package searchelement

func SearchElement(A []int, B int) bool {

	for i := 0; i < len(A); i++ {
		if A[i] == B {
			return true
		}
	}
	return false
}
