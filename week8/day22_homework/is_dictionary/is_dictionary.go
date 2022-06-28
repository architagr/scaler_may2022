package is_dictionary

func IsDisctionary(A []string, B string) int {

	letterMap := convertLettersToMap(B)
	n := len(A)
	k := 0
	for i := 1; i < n; i, k = i+1, 0 {
		currentStr, oldString := A[i], A[i-1]
		if currentStr[k] == oldString[k] {
			min := len(currentStr)
			if min > len(oldString) {
				min = len(oldString)
			}
			count := 1
			k++
			for ; k < min; k = k + 1 {
				if currentStr[k] != oldString[k] {
					if letterMap[currentStr[k]] < letterMap[oldString[k]] {
						return 0
					} else {
						//practically breaking out of loop of k
						k = min
					}

				} else {
					// if complete string matches like hell and hello so hell should come first
					count++
				}
			}
			if count == min && len(oldString) != min {
				return 0
			}

		} else if letterMap[currentStr[k]] < letterMap[oldString[k]] {
			return 0
		}
	}
	return 1
}
func convertLettersToMap(b string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < 26; i++ {
		m[b[i]] = i
	}
	return m
}
