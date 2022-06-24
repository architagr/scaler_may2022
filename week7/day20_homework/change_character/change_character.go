package change_character

func ChangeCharacter(A string, B int) int {
	charArray := make([]int, 26)
	max := -1
	n := len(A)
	for i := 0; i < n; i++ {
		x := A[i] - 'a'

		charArray[int(x)]++
		if charArray[int(x)] > max {
			max = charArray[int(x)]
		}
	}

	for i := 1; B > 0 && B >= i && i < max; i++ {
		for j := 0; j < 26; j++ {
			if charArray[j] > 0 && charArray[j] <= B && charArray[j] == i {
				charArray[j] = 0
				B -= i
			}
		}
	}
	count := 0
	for j := 0; j < 26; j++ {
		if charArray[j] > 0 {
			count++
		}
	}
	return count
}
