package invert_the_binary_tree

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
 * @Output root pointer of tree.
 */
func InvertTree(A *treeNode) *treeNode {
	if A == nil {
		return A
	}
	A.left = InvertTree(A.left)
	A.right = InvertTree(A.right)
	A.left, A.right = A.right, A.left
	return A
}
