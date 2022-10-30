package identical_binary_trees

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
 * @input B : Root pointer of the tree
 *
 * @Output Integer
 */
func IsSameTree(A *treeNode, B *treeNode) int {
	if A == nil && B == nil {
		return 1
	} else if A == nil || B == nil {
		return 0
	} else if A.value != B.value {
		return 0
	}
	x := IsSameTree(A.left, B.left)
	y := IsSameTree(A.right, B.right)

	if x == 0 || y == 0 {
		return 0
	}
	return 1
}
