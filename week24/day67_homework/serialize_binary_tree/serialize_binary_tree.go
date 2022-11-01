package serialize_binary_tree

import "fmt"

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

type Node struct {
	Data *treeNode
	next *Node
}

func NewNode(value *treeNode) *Node {
	newNode := new(Node)
	newNode.Data = value
	return newNode
}

type Queue struct {
	count      int
	head, tail *Node
}

type IQueue interface {
	Enqueue(data *treeNode)
	Dequeue() *Node
	IsEmpty() bool
	Peep() (result *treeNode, err error)
}

func NewQueue() IQueue {
	return new(Queue)
}
func (queue *Queue) Enqueue(data *treeNode) {
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
func (queue *Queue) Peep() (result *treeNode, err error) {
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

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer array.
 */
func SerializeBinaryTree(A *treeNode) []int {
	if A == nil {
		return []int{}
	}
	arr := make([]int, 0)
	st := NewQueue()
	st.Enqueue(A)
	for !st.IsEmpty() {
		n := st.Dequeue()
		arr = append(arr, n.Data.value)
		if n.Data.value != -1 {
			left := treeNode_new(-1)
			if n.Data.left != nil {
				left = n.Data.left
			}

			right := treeNode_new(-1)
			if n.Data.right != nil {
				right = n.Data.right
			}

			st.Enqueue(left)
			st.Enqueue(right)
		}
	}
	return arr
}
