package node_count

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func CountNode(A *treeNode) int {

	if A == nil {
		return 0
	}

	return 1 + CountNode(A.left) + CountNode(A.right)
}
