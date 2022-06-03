package multiple_left_rotation_of_array

import (
	reversearray "scaler-may-22/week3/day8_assingment/reverse_array"
)

func MultipleLeftRotateArray(A []int, B []int) [][]int {
	result := make([][]int, 0)
	for _, val := range B {
		x := LeftRotateArray(A, val)
		result = append(result, x)
	}
	return result
}

func LeftRotateArray(A []int, k int) []int {
	k = k % len(A)
	result := make([]int, len(A))
	copy(result, A)

	result = reversearray.ReverseArray(result, 0, len(A)-1)
	result = reversearray.ReverseArray(result, len(A)-k, len(A)-1)
	result = reversearray.ReverseArray(result, 0, len(A)-k-1)
	return result
}
