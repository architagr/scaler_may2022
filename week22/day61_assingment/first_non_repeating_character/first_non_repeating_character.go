package first_non_repeating_character

import "fmt"

type Node struct {
	data byte
	next *Node
}

type Queue struct {
	count      int
	head, tail *Node
}

type IQueue interface {
	Enqueue(data byte)
	Dequeue() *Node
	IsEmpty() bool
}

func (queue *Queue) Enqueue(data byte) {
	node := new(Node)
	node.data = data
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
func (queue *Queue) Peep() (byte, error) {
	if queue.IsEmpty() {
		return '#', fmt.Errorf("no data")
	}
	return queue.head.data, nil
}

func (queue *Queue) IsEmpty() bool {
	return queue.count == 0
}

func FirstNonRepeatingChar(A string) string {
	hashMap := make(map[byte]int)
	queue := new(Queue)
	result := make([]byte, len(A))
	hashMap[A[0]]++
	queue.Enqueue(A[0])
	result[0] = A[0]
	for i := 1; i < len(A); i++ {
		hashMap[A[i]]++
		queue.Enqueue(A[i])
		x, err := queue.Peep()
		for ; err == nil && hashMap[x] > 1; x, err = queue.Peep() {
			queue.Dequeue()
		}
		result[i] = x
	}
	return string(result)
}
