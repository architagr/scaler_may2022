package first_depth_first_search

import "fmt"
func FirstDepthFirstSearch(A []int , B int , C int )  (int) {
     visited := make([]bool, len(A)+1)
	queue := NewQueue()
	at := make(map[int]map[int]bool)
	for i := 0; i < len(A); i++ {
		val, ok := at[A[i]]
		if !ok {
			val = make(map[int]bool)
		}
		val[i+1] = true
		at[A[i]] = val
	}
	queue.Enqueue(C)
	visited[C] = true
	for !queue.IsEmpty() {
		x := queue.Dequeue().Data
		val := at[x]
		for key := range val {
			if !visited[key] {
				queue.Enqueue(key)
				visited[key] = true
			}
		}
	}
	if visited[B] {
		return 1
	}
	return 0
}


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