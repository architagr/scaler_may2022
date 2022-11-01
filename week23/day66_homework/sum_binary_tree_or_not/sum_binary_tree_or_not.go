package sum_binary_tree_or_not

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
 * @Output Integer
 */
func IsSumBinaryTree(A *treeNode) int {

	_, ok := sumCheck(A)
	return ok
}

func sumCheck(A *treeNode) (sum int, ok int) {
	if A == nil {
		return 0, 1
	}
	if A.left == nil && A.right == nil {
		return A.value, 1
	}

	left, okL := sumCheck(A.left)
	right, okr := sumCheck(A.right)

	if okL == 1 && okr == 1 && (left+right) == A.value {
		return (left + right + A.value), 1
	}
	return 0, 0
}
