package next_pointer_binary_tree

import "fmt"

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
	next  *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	node.next = nil
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
	Count() int
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

func (queue *Queue) Count() int {
	return queue.count
}

func Connect(root *treeNode) *treeNode {

	if root == nil {
		return root
	}
	queue := NewQueue()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		var prev *treeNode = nil
		size := queue.Count()
		for i := 0; i < size; i++ {
			tmp := queue.Dequeue().Data
			if tmp.left != nil {
				queue.Enqueue(tmp.left)
				if prev != nil {
					prev.next = tmp.left
				}
			}
			if tmp.right != nil {
				queue.Enqueue(tmp.right)
				tmp.left.next = tmp.right
				if i == size-1 {
					tmp.right.next = nil
				} else {
					prev = tmp.right
				}
			}

		}
	}
	return root
}
