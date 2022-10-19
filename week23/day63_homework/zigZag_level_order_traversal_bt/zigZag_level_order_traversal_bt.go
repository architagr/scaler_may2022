package zigZag_level_order_traversal_bt

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
 * @Output 2D int array.
 */
func ZigzagLevelOrder(A *treeNode) [][]int {
	result := make([][]int, 0)
	parse(A, &result, 0)
	for i := 1; i < len(result); i = i + 2 {
		l, r := 0, len(result[i])-1
		for l <= r {
			result[i][l], result[i][r] = result[i][r], result[i][l]
			l++
			r--
		}
	}
	return result
}

func parse(head *treeNode, result *[][]int, level int) {
	if head == nil {
		return
	}
	if len(*result)-1 < level {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], head.value)

	parse(head.left, result, level+1)
	parse(head.right, result, level+1)

}
