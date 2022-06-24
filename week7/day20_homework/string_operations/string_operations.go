package string_operations

func StringOperation(A string) string {
	k, n := 0, len(A)
	str := []byte(A)
	for i := 0; i < n; i++ {
		if checkCapitalLetter(str[i]) {
			if i < n-1 {
				str[k] = str[i+1]
			}
		} else {
			if checkVowel(str[i]) {
				str[k] = '#'
			} else {
				str[k] = str[i]
			}

			k++
		}
	}

	return string(append(str[:k], str[:k]...))
}

func checkVowel(a byte) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u'
}

func checkCapitalLetter(a byte) bool {
	return a >= 'A' && a <= 'Z'
}
