package reverse_string_by_word

import (
	"scaler-may-22/week7/day20_assingment/reverse_string"
	"strings"
)

func ReverseStringByWord(A string) string {
	A = strings.TrimSpace(A)
	A = string(removeMultipleSpaces([]byte(A)))
	A = string(reverse_string.ReverseStringPart([]byte(A), 0, len(A)-1))
	lastSpace := -1
	for i := 0; i < len(A); i++ {
		if A[i] == ' ' {
			A = string(reverse_string.ReverseStringPart([]byte(A), lastSpace+1, i-1))
			lastSpace = i
		}
	}
	A = string(reverse_string.ReverseStringPart([]byte(A), lastSpace+1, len(A)-1))
	return A
}

func removeMultipleSpaces(A []byte) []byte {
	for i := 1; i < len(A); i++ {
		if A[i] == ' ' && A[i-1] == ' ' {
			A = append(A[0:i], A[i+1:]...)
		}
	}
	return A
}
