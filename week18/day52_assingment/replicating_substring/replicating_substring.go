package replicating_substring

func ReplicatingSubstring(A int, B string) int {
	hashmap := make(map[byte]int)
	for i := 0; i < len(B); i++ {
		hashmap[B[i]]++
	}
	for _, count := range hashmap {
		if count%A != 0 {
			return -1
		}
	}
	return 1
}
