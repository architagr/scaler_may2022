package reverse_string

func ReverseString(A string) string {
	return string(ReverseStringPart([]byte(A), 0, len(A)-1))
}

func ReverseStringPart(A []byte, start, end int) []byte {
	for start <= end {
		temp := A[start]
		A[start] = A[end]
		A[end] = temp
		start++
		end--
	}

	return A
}
