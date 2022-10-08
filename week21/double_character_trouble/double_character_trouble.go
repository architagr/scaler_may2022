package double_character_trouble

import "scaler-may-22/common"

func DoubleCharTrouble(A string) string {
	stack := new(common.Stack[byte])
	stack.Push(A[0])
	for i := 1; i < len(A); i++ {
		top := stack.Top()
		if top == A[i] {
			stack.Pop()
		} else {
			stack.Push(A[i])
		}
	}
	result := make([]byte, 0)

	for !stack.IsEmpty() {
		result = append(result, stack.Pop().Data)
	}
	l, r := 0, len(result)-1

	for l <= r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}
	return string(result)
}
