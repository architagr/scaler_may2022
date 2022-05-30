package amazing_subarray

func FindCountOfAmazingSubarray(A string) int {
	n := len(A)
	count := 0

	for i, cha := range A {
		str := string(cha)
		if str == "A" || str == "a" || str == "E" || str == "e" || str == "I" || str == "i" || str == "O" || str == "o" || str == "U" || str == "u" {
			count += n - i
		}
	}
	return count % 10003
}
