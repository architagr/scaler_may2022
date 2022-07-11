package inorder_traversal

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer array.
 */
func InorderTraversal(A *treeNode) []int {
	arr := make([]int, 0)
	return inorderTraversal(A, arr)
}

func inorderTraversal(A *treeNode, arr []int) []int {
	if A == nil {
		return arr
	}

	arr = inorderTraversal(A.left, arr)
	arr = append(arr, A.value)
	arr = inorderTraversal(A.right, arr)
	return arr
}
