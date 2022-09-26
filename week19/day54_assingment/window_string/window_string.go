package window_string

import "math"

type SubString struct {
	start, end, length int
}

func MinWindowString(A, B string) string {
	answer := SubString{
		start:  0,
		end:    -1,
		length: math.MaxInt,
	}
	size_A, count, total_Char := len(A), 0, len(B)
	// FILL HASHMAP for B
	hashMap := fill_Map(B)

	for i, j := 0, 0; j < size_A; j++ {
		current := A[j]
		// IF THE CHARACTER IS PRESENT IN B, INSERT IT
		if _, ok := hashMap[current]; ok {
			count = insert_Current_Character(A, j, count, hashMap)
		}
		// REMOVE THE EXTRA CHARACTERS WHICH IS EXTRA IN CURRENT WINDOW
		i = remove_Extra_Front_Chars(A, i, j, hashMap)
		// IF CURRENT COUNT == TOTAL REQUIRED CHARS, AND CURRENT WONDOW LENGTH IS < PRESENT LENGTH
		if count == total_Char && (j-i+1) < answer.length {
			answer = update_Length(i, j, answer)
		}
	}

	sb := ""
	if answer.length < math.MaxInt {
		sb = A[answer.start : answer.end+1]
	}

	return sb
}

func insert_Current_Character(str string, index, count int, hashMap map[byte]int) int {
	current := str[index]
	if freq, ok := hashMap[current]; ok {
		// WE DECREASE THE FREQ AS THIS FREQ INDICATES HOW MANY MORE OF CURRENT CHARACTER IS NEEDED
		freq--
		hashMap[current] = freq
		// IF NEW FREQ == 0, IT MEANS WE HAVE EXACTLY SAME NUMBER OF CURRENT CHARACTER NEEDED, INCREASE COUNT
		if freq >= 0 {
			count++
		}
	}
	return count
}

func update_Length(start, end int, answer SubString) SubString {
	answer.start = start
	answer.end = end
	answer.length = end - start + 1
	return answer
}

func remove_Extra_Front_Chars(str string, start, end int, hashMap map[byte]int) int {
	for start <= end {
		char_At_First := str[start]
		if freq_At_Start, ok := hashMap[char_At_First]; ok {
			if freq_At_Start < 0 {
				hashMap[char_At_First]++
				start++
			} else {
				break
			}
		} else {
			start++
		}
	}
	return start
}

func fill_Map(str string) map[byte]int {
	size := len(str)
	hashMap := make(map[byte]int)
	for i := 0; i < size; i++ {
		current := str[i]
		hashMap[current]++
	}
	return hashMap
}
