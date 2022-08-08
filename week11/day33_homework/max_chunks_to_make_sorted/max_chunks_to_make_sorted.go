package max_chunks_to_make_sorted

func MaxChunk(A []int) int {
	chunks := 0
	curr := -1
	for i := 0; i < len(A); i++ {
		curr = maxvalue(curr, A[i])
		if curr == i {
			chunks++
		}
	}
	return chunks
}

func maxvalue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
