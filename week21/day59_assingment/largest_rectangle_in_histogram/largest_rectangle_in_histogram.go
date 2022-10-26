package largest_rectangle_in_histogram

import "math"

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
func (stack *Stack) Top() int {
	if stack.head == nil {
		return -1
	}
	return stack.head.val
}
func prevSmaller(A []int) []int {
	result := make([]int, len(A))

	stack := Stack{
		count: 0,
		head:  nil,
	}
	for i := 0; i < len(A); i++ {
		for stack.head != nil && A[stack.Top()] >= A[i] {
			stack.Pop()
		}
		if stack.head == nil {
			result[i] = -1
		} else {
			result[i] = stack.Top()
		}
		stack.Push(i)
	}
	return result
}

func nextSmaller(A []int) []int {
	result := make([]int, len(A))

	stack := Stack{
		count: 0,
		head:  nil,
	}
	for i := len(A) - 1; i > -1; i-- {
		for stack.head != nil && A[stack.Top()] >= A[i] {
			stack.Pop()
		}
		if stack.head == nil {
			result[i] = len(A)
		} else {
			result[i] = stack.Top()
		}
		stack.Push(i)
	}
	return result
}

func LargestRectangleArea(A []int) int {
	n := len(A)
	if n == 1 {
		return A[0]
	}
	SL := prevSmaller(A)
	SR := nextSmaller(A)

	largeArea := math.MinInt

	for i := 0; i < n; i++ {
		area := A[i] * (SR[i] - SL[i] - 1)
		largeArea = maxValue(largeArea, area)
	}
	return largeArea
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
