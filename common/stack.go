package common

import "fmt"

type Node[T any] struct {
	Data T
	Prev *Node[T]
}
type Stack[T any] struct {
	count int
	Head  *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	newNode := new(Node[T])
	newNode.Data = value
	return newNode
}

func (stack *Stack[T]) Push(value T) {
	newNode := NewNode(value)
	stack.count++
	if stack.Head == nil {
		stack.Head = newNode
		return
	}
	newNode.Prev = stack.Head
	stack.Head = newNode

}

func (stack *Stack[T]) Pop() *Node[T] {
	if stack.Head == nil {
		return nil
	}
	node := stack.Head
	stack.Head = stack.Head.Prev
	stack.count--
	return node
}

func (stack *Stack[T]) Top() (value T) {
	if stack.Head == nil {
		return
	}
	value = stack.Head.Data
	return
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.count == 0
}

func (stack *Stack[T]) Count() int {
	fmt.Fprint()
	return stack.count
}
