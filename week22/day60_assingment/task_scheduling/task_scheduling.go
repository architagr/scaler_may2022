package task_scheduling

func TaskSchedule(A []int, B []int) int {
	count, headIndex := 0, 0

	for _, task := range B {
		for ; A[headIndex] != task; headIndex++ {
			A = append(A, A[headIndex])
			count++
		}
		headIndex++
		count++
	}
	return count
}
