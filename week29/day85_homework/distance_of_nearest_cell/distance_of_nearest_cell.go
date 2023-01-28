package distance_of_nearest_cell

import (
	"fmt"
	"math"
)

type Pair struct {
	x, y int
}
type Node struct {
	Data Pair
	next *Node
}

func NewNode(value Pair) *Node {
	newNode := new(Node)
	newNode.Data = value
	return newNode
}

type Queue struct {
	count      int
	head, tail *Node
}

type IQueue interface {
	Enqueue(data Pair)
	Dequeue() *Node
	IsEmpty() bool
	Peep() (result Pair, err error)
}

func NewQueue() IQueue {
	return new(Queue)
}
func (queue *Queue) Enqueue(data Pair) {
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
func (queue *Queue) Peep() (result Pair, err error) {
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

var que *Queue

func solve(A [][]int) [][]int {
	que = new(Queue)
	n, m := len(A), len(A[0])
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				que.Enqueue(Pair{x: i, y: j})
			} else {
				mat[i][j] = math.MaxInt
			}
		}
	}

	return bfs(A, mat)
}

func bfs(A, mat [][]int) [][]int {
	direction := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for !que.IsEmpty() {
		pair := que.Dequeue()
		for i := 0; i < len(direction); i++ {
			new_x := (*pair).Data.x + direction[i][0]
			new_y := (*pair).Data.y + direction[i][1]
			if getDirection(new_x, new_y, A) {
				if mat[new_x][new_y] > mat[(*pair).Data.x][(*pair).Data.y] {
					mat[new_x][new_y] = mat[(*pair).Data.x][(*pair).Data.y] + 1
					que.Enqueue(Pair{x: new_x, y: new_y})
				}
			}
		}
	}
	return mat
}
func getDirection(x, y int, matrix [][]int) bool {
	return x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0])
}
