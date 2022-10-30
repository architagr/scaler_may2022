package maximum_frequency_stack

import "scaler-may-22/common"

func MaxFrequencyStack(A [][]int) []int {
	result := make([]int, len(A))

	countHm := make(map[int]int)
	frequencyHm := make(map[int]*common.Stack[int])
	highestFrequency := 0
	for i := 0; i < len(A); i++ {
		if A[i][0] == 1 {
			countHm[A[i][1]]++
			freq := countHm[A[i][1]]

			if freq > highestFrequency {
				highestFrequency = freq
			}

			st := frequencyHm[freq]
			if st == nil {
				st = new(common.Stack[int]) //for first time
				st.Push(A[i][1])
			} else {
				st.Push(A[i][1])
			}
			frequencyHm[freq] = st //Maintain the stack for each frequency
			result[i] = -1         //add to ans
		} else {
			st := frequencyHm[highestFrequency] //get highestFrequency stack
			val := st.Pop().Data

			countHm[val]--

			if st.IsEmpty() { //if stack is isEmpty reduce highestFrequency count to prevoius count
				highestFrequency--
			}

			result[i] = val //add to ans
		}
	}
	return result
}
