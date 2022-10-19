package balanced_binary_tree

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
func isBalanced(A *treeNode) int {
	_, ok := treeHeight(A)
	if ok {
		return 1
	}
	return 0
}

func treeHeight(head *treeNode) (height int, ok bool) {

	if head == nil || head.value == -1 {
		height = 0
		ok = true
		return
	}
	if head.left == nil && head.right == nil {
		height = 1
		ok = true
		return
	}
	lHeight, rHeight := 0, 0
	left, right := false, false

	lHeight, left = treeHeight(head.left)
	rHeight, right = treeHeight(head.right)
	diff := (lHeight - rHeight)
	if diff < 0 {
		diff *= -1
	}
	return maxValue(lHeight, rHeight) + 1, left && right && diff <= 1
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
