package symmetric_binary_tree

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
func isSymmetric(A *treeNode) int {
	return isSame(A.left, invertTree(A.right))
}

func isSame(A, B *treeNode) int {
	if A == nil && B == nil {
		return 1
	} else if A == nil || B == nil {
		return 0
	}
	if A.value != B.value {
		return 0
	}
	left := isSame(A.left, B.left)
	right := isSame(A.right, B.right)

	if left == 1 && right == 1 {
		return 1
	}
	return 0
}

func invertTree(A *treeNode) *treeNode {
	if A == nil {
		return A
	}
	A.left = invertTree(A.left)
	A.right = invertTree(A.right)
	A.left, A.right = A.right, A.left
	return A
}
