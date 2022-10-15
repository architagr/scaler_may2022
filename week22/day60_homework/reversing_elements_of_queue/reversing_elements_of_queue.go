package reversing_elements_of_queue

import "scaler-may-22/common"

type Node struct {
	Data int
	next *Node
}

type Queue struct {
	head, tail *Node
}

func (q *Queue) Enqueue(data int) {
	node := new(Node)
	node.Data = data

	if q.head == nil && q.tail == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}
func (q *Queue) Dequeue() (result *Node) {
	result = q.head
	if q.head == nil {
		return
	}
	q.head = q.head.next
	return
}
func ReverseElementsOfQueue(A []int, B int) []int {
	queue := new(Queue)
	stack := new(common.Stack[int])
	i := 0
	for ; i < B; i++ {
		queue.Enqueue(A[i])
		stack.Push(A[i])
	}
	for ; i < len(A); i++ {
		queue.Enqueue(A[i])
	}

	result := make([]int, 0, len(A))
	for !stack.IsEmpty() {
		result = append(result, stack.Pop().Data)
		queue.Dequeue()
	}
	for len(result) < len(A) {
		result = append(result, queue.Dequeue().Data)
	}
	return result
}
