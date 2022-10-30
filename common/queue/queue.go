package common

import "fmt"

type Node[T any] struct {
	Data T
	next *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	newNode := new(Node[T])
	newNode.Data = value
	return newNode
}

type Queue[T any] struct {
	count      int
	head, tail *Node[T]
}

type IQueue[T any] interface {
	Enqueue(data T)
	Dequeue() *Node[T]
	IsEmpty() bool
	Peep() (result T, err error)
}

func NewQueue[T any]() IQueue[T] {
	return new(Queue[T])
}
func (queue *Queue[T]) Enqueue(data T) {
	node := NewNode(data)
	if queue.IsEmpty() {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
	queue.count++
}

func (queue *Queue[T]) Dequeue() *Node[T] {
	if queue.IsEmpty() {
		return nil
	}
	node := queue.head
	queue.head = queue.head.next
	queue.count--
	if queue.IsEmpty() {
		queue.head = nil
		queue.tail = nil
	}
	return node
}
func (queue *Queue[T]) Peep() (result T, err error) {
	if queue.IsEmpty() {
		err = fmt.Errorf("no data")
		return
	}
	result = queue.head.Data
	return
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.count == 0
}
