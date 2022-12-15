package rotten_oranges

import "fmt"

type Point struct {
	i, j int
}
type Node struct {
	Data Point
	next *Node
}

func NewNode(value Point) *Node {
	newNode := new(Node)
	newNode.Data = value
	return newNode
}

type Queue struct {
	count      int
	head, tail *Node
}

type IQueue interface {
	Enqueue(data Point)
	Dequeue() *Node
	IsEmpty() bool
	Peep() (result Point, err error)
}

func NewQueue() IQueue {
	return new(Queue)
}
func (queue *Queue) Enqueue(data Point) {
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
func (queue *Queue) Peep() (result Point, err error) {
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

var queue IQueue
var visited [][]bool

func RottenOranges(A [][]int) int {
	count := 0
	n, m := len(A), len(A[0])
	queue = NewQueue()
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 2 {
				validate(A, i, j)
			}
		}
	}
	if !queue.IsEmpty() {
		queue.Enqueue(Point{i: -1, j: -1})
		for !queue.IsEmpty() {
			point := queue.Dequeue().Data
			if point.i == -1 && point.j == -1 {
				count++
				if !queue.IsEmpty() {
					queue.Enqueue(Point{i: -1, j: -1})
				}
				continue
			}
			i := point.i
			j := point.j
			A[i][j] = 2
			validate(A, i, j)
		}
	}

	if check(A) > 0 {
		return -1
	}
	return count
}
func validate(A [][]int, i, j int) {
	n, m := len(A), len(A[0])
	visited[i][j] = true
	if i+1 < n && !visited[i+1][j] && A[i+1][j] == 1 {
		queue.Enqueue(Point{i: i + 1, j: j})
		visited[i+1][j] = true
	}
	if i-1 >= 0 && !visited[i-1][j] && A[i-1][j] == 1 {
		queue.Enqueue(Point{i: i - 1, j: j})
		visited[i-1][j] = true
	}
	if j+1 < m && !visited[i][j+1] && A[i][j+1] == 1 {
		queue.Enqueue(Point{i: i, j: j + 1})
		visited[i][j+1] = true
	}
	if j-1 >= 0 && !visited[i][j-1] && A[i][j-1] == 1 {
		queue.Enqueue(Point{i: i, j: j - 1})
		visited[i][j-1] = true
	}
}
func check(A [][]int) int {
	count := 0
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				count++
			}
		}
	}
	return count

}
