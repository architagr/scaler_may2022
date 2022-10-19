package odd_and_even_levels

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
func OddEvenDiff(A *treeNode) int {
	result := make([]int, 2)
	parse(A, &result, 0)
	return result[0] - result[1]
}

func parse(head *treeNode, result *[]int, level int) {
	if head == nil {
		return
	}

	(*result)[level%2] += head.value
	parse(head.left, result, level+1)
	parse(head.right, result, level+1)
}
