package cycle_in_undirected_graph

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

func solve(A int, B [][]int) int {
	visited := make([]bool, A+1)
	count := 0

	at := make(map[int]map[int]bool)
	for i := 0; i < len(B); i++ {
		val, ok := at[B[i][0]]
		if !ok {
			val = make(map[int]bool)
		}
		val[B[i][1]] = true
		at[B[i][0]] = val

		val, ok = at[B[i][1]]
		if !ok {
			val = make(map[int]bool)
		}
		val[B[i][0]] = true
		at[B[i][1]] = val
	}
	queue := NewQueue()
	for i := 1; i <= A; i++ {
		if !visited[i] {
			count++
			queue.Enqueue(i)
			visited[i] = true
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
		}
	}

	if len(B) > A-count {
		return 1
	}
	return 0
}
