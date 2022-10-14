package level_order

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
 * @Output 2D int array.
 */
func levelOrder(A *treeNode) [][]int {
	result := make([][]int, 0)
	parse(A, &result, 0)
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
