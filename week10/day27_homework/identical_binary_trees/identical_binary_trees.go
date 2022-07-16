package identical_binary_trees

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func isSameTree(A *treeNode, B *treeNode) int {
	if A == nil && B == nil {
		return 1
	} else if (A != nil && B == nil) || (B != nil && A == nil) {
		return 0
	}
	if A.value != B.value {
		return 0
	}
	left := isSameTree(A.left, B.left)
	if left == 0 {
		return 0
	}
	right := isSameTree(A.right, B.right)
	if right == 0 {
		return 0
	}
	return 1
}
