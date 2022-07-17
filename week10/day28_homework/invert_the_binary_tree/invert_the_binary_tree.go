package invert_the_binary_tree

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func InvertTree(A *treeNode) *treeNode {
	if A == nil {
		return nil
	}
	A.right, A.left = A.left, A.right
	A.left = InvertTree(A.left)
	A.right = InvertTree(A.right)
	return A
}
