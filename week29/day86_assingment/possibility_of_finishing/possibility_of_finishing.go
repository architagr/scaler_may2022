package possibility_of_finishing

import "fmt"

type Node struct {
	Data int
	next *Node
}

func NewNode(value int) *Node {
	newNode := new(Node)
	newNode.Data = value
	return newNode
}

type Queue struct {
	count      int
	head, tail *Node
}

type IQueue interface {
	Enqueue(data int)
	Dequeue() *Node
	IsEmpty() bool
	Peep() (result int, err error)
}

func NewQueue() IQueue {
	return new(Queue)
}
func (queue *Queue) Enqueue(data int) {
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

func (queue *Queue) Dequeue() *Node {
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
func (queue *Queue) Peep() (result int, err error) {
	if queue.IsEmpty() {
		err = fmt.Errorf("no data")
		return
	}
	result = queue.head.Data
	return
}

func (queue *Queue) IsEmpty() bool {
	return queue.count == 0
}

func PosibilityOfFinishing(A int, B, C []int) int {
	visited := make([]int, A+1)
	at := make(map[int]map[int]bool)
	for i := 0; i < len(B); i++ {
		val, ok := at[B[i]]
		if !ok {
			val = make(map[int]bool)
		}
		if _, ok := val[C[i]]; !ok {
			val[C[i]] = true
			at[B[i]] = val
			visited[C[i]]++
		}
	}
	queue := NewQueue()
	for i := 1; i < len(visited); i++ {
		if visited[i] == 0 {
			queue.Enqueue(i)
		}
	}
	count := 0
	for !queue.IsEmpty() {
		min := queue.Dequeue().Data
		count++
		for k := range at[min] {
			visited[k]--
			if visited[k] == 0 {
				queue.Enqueue(k)
			}
		}
	}
	if count < A {
		return 0
	}
	return 1
}
