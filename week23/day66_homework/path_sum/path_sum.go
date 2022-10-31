package path_sum

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
 * @input B : Integer
 *
 * @Output Integer
 */
func HasPathSum(root *treeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := sum(root, targetSum, 0)
	if ans {
		return 1
	} else {
		return 0
	}
}

func sum(root *treeNode, target, current int) bool {
	if root == nil {
		return false
	} else if root.left != nil && root.right != nil {
		return sum(root.left, target, current+root.value) || sum(root.right, target, current+root.value)
	} else if root.left != nil {
		return sum(root.left, target, current+root.value)
	} else if root.right != nil {
		return sum(root.right, target, current+root.value)
	}
	return target == root.value+current
}
