package permutations_of_a_in_b

import "reflect"

func PermutationAinB(A string, B string) int {
	charCountA := make([]int, 26)
	charCountB := make([]int, 26)
	for i := 0; i < len(A); i++ {
		charCountA[A[i]-'a'] += 1
	}
	for i := 0; i < len(A); i++ {
		charCountB[B[i]-'a'] += 1
	}
	count := 0
	if reflect.DeepEqual(charCountA, charCountB) {
		count++
	}
	start, end := 1, len(A)
	for end < len(B) {
		charCountB[B[start-1]-'a'] -= 1
		charCountB[B[end]-'a'] += 1
		if reflect.DeepEqual(charCountA, charCountB) {
			count++
		}
		start++
		end++
	}
	return count
}
