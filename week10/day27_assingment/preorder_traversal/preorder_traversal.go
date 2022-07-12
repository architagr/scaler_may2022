package preorder_traversal

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func PreorderTraversal(A *treeNode) []int {

	arr := make([]int, 0)
	return traversal(A, arr)
}

func traversal(A *treeNode, arr []int) []int {
	if A == nil {
		return arr
	}
	arr = append(arr, A.value)
	arr = traversal(A.left, arr)

	arr = traversal(A.right, arr)
	return arr
}
