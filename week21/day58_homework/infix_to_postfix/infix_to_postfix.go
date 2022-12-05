package infix_to_postfix

import "scaler-may-22/common"

func InfixToPostfix(infix string) string {

	stack := new(common.Stack[byte])
	postfix := make([]byte, 0, len(infix))
	for i := 0; i < len(infix); i++ {
		if infix[i] <= 'z' && infix[i] >= 'a' {
			postfix = append(postfix, infix[i])
		} else if stack.IsEmpty() || infix[i] == '(' || Precedence(stack.Top()) < Precedence(infix[i]) {
			stack.Push(infix[i])
		} else if infix[i] == ')' {
			for !stack.IsEmpty() && stack.Top() != '(' {
				n := stack.Pop()
				postfix = append(postfix, n.Data)
			}
			stack.Pop()
		} else {
			for !stack.IsEmpty() && Precedence(stack.Top()) >= Precedence(infix[i]) {
				no := stack.Pop()
				postfix = append(postfix, no.Data)
			}
			stack.Push(infix[i])
		}
	}
	for !stack.IsEmpty() {
		no := stack.Pop()
		postfix = append(postfix, no.Data)
	}
	return string(postfix)
}

func Precedence(A byte) int {
	if A == '^' {
		return 3
	} else if A == '*' || A == '/' {
		return 2
	} else if A == '(' {
		return 0
	}
	return 1
}
