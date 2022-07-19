package identical_binary_trees

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func isSameTree(A *treeNode, B *treeNode) int {
	if A == nil && B == nil {
		return 1
	} else if A == nil || B == nil {
		return 0
	} else if A.value != B.value {
		return 0
	}
	if isSameTree(A.left, B.left) == 0 || isSameTree(A.right, B.right) == 0 {
		return 0
	}

	return 1

}
