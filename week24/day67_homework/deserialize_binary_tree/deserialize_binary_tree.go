package deserialize_binary_tree

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
 * @input A : Integer array
 *
 * @Output root pointer of tree.
 */
func DeserializeBinaryTree(A []int) *treeNode {
	if len(A) == 0 {
		return nil
	}

	queue := NewQueue()
	root := treeNode_new(A[0])
	queue.Enqueue(root)
	for i := 1; i < len(A); i = i + 2 {
		node := queue.Dequeue()
		if A[i] != -1 {
			n1 := treeNode_new(A[i])
			node.Data.left = n1
			queue.Enqueue(n1)
		}
		if A[i+1] != -1 {
			n1 := treeNode_new(A[i+1])
			node.Data.right = n1
			queue.Enqueue(n1)
		}
	}
	return root
}
