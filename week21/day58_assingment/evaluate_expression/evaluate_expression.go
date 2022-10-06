package evaluate_expression

import "strconv"

type Node struct {
	val  int
	prev *Node
}

func NewNode(value int) *Node {
	newNode := new(Node)
	newNode.val = value
	return newNode
}

type Stack struct {
	count int
	head  *Node
}

func (stack *Stack) Push(value int) {

	newNode := NewNode(value)
	stack.count++
	if stack.head == nil {
		stack.head = newNode
		return
	}
	newNode.prev = stack.head
	stack.head = newNode

}

func (stack *Stack) Pop() *Node {
	if stack.head == nil {
		return nil
	}
	node := stack.head
	stack.head = stack.head.prev
	stack.count--
	return node
}

func EvalRPN(A []string) int {
	if len(A) == 0 {
		return 0
	}
	var stack = new(Stack)
	for i := 0; i < len(A); i++ {
		val := A[i]
		switch val {
		case "+":
			a1 := stack.Pop().val
			a2 := stack.Pop().val
			stack.Push(a2 + a1)
		case "-":
			a1 := stack.Pop().val
			a2 := stack.Pop().val
			stack.Push(a2 - a1)
		case "*":
			a1 := stack.Pop().val
			a2 := stack.Pop().val
			stack.Push(a2 * a1)
		case "/":
			a1 := stack.Pop().val
			a2 := stack.Pop().val
			stack.Push(a2 / a1)
		default:
			x, _ := strconv.Atoi(val)
			stack.Push(x)
		}
	}

	return stack.Pop().val
}
