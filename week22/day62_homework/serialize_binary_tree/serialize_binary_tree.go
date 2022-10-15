package serialize_binary_tree

type Node struct {
	Data *treeNode
	next *Node
}

type Queue struct {
	head, tail *Node
}

func (q *Queue) Enqueue(data *treeNode) {
	node := new(Node)
	node.Data = data

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
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return
	}
	q.head = q.head.next
	return
}
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

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

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer array.
 */
func SerializeBinaryTree(A *treeNode) []int {
	result := make([]int, 0)
	queue := new(Queue)
	queue.Enqueue(A)
	for !queue.IsEmpty() {
		top := queue.Dequeue()
		if top.Data == nil {
			result = append(result, -1)
		} else {
			result = append(result, top.Data.value)

			queue.Enqueue(top.Data.left)
			queue.Enqueue(top.Data.right)
		}
	}
	return result
}
