package tree_height

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func CalcHeight(head *treeNode) int {

	if head == nil {
		return 0
	}
	right := CalcHeight(head.right) + 1
	left := CalcHeight(head.left) + 1

	if left > right {
		return left
	}
	return right
}
