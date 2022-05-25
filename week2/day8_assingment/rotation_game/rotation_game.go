package rotationgame

import (
	reversearray "scaler-may-22/week2/day8_assingment/reverse_array"
)

func RotateArray(arr []int, k int) []int {
	k = k % len(arr)
	arr = reversearray.ReverseArray(arr, 0, len(arr)-1)
	arr = reversearray.ReverseArray(arr, 0, k-1)
	arr = reversearray.ReverseArray(arr, k, len(arr)-1)
	return arr
}
