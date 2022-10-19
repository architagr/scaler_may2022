package top_view_of_binary_tree

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
	node  *treeNode
	index int
	next  *Node
}

func NewNode(value *treeNode, index int) *Node {
	var newNode = &Node{
		node:  value,
		index: index,
	}
	return newNode
}

type Queue struct {
	head, tail *Node
	count      int
}

func (q *Queue) Enqueue(value *treeNode, index int) {
	node := NewNode(value, index)
	q.count++
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
	q.count--
	if q.count == 0 {
		q.head = nil
		q.tail = nil
	}
	return
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer array.
 */

func TopVeiwOfBTree(A *treeNode) []int {
	queue := new(Queue)
	hashMap := make(map[int]int)
	max, min := 0, 0

	queue.Enqueue(A, 0)
	queue.Enqueue(nil, 0)
	for queue.count > 0 {
		node := queue.Dequeue()
		if node.node != nil {
			if _, ok := hashMap[node.index]; !ok {
				hashMap[node.index] = node.node.value
			}
			if node.index > max {
				max = node.index
			}
			if node.index < min {
				min = node.index
			}
			if node.node.left != nil && node.node.left.value != -1 {
				queue.Enqueue(node.node.left, node.index-1)
			}
			if node.node.right != nil && node.node.right.value != -1 {
				queue.Enqueue(node.node.right, node.index+1)
			}
		} else if queue.count > 0 {
			queue.Enqueue(nil, 0)
		}
	}
	result := make([]int, 0)
	for i := min; i <= max; i++ {
		result = append(result, hashMap[i])
	}
	return result
}
