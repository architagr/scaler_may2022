package symmetric_binary_tree

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func IsSymmetric(A *treeNode) int {
	inorder := InorderTraversal(A)
	if len(inorder)%2 == 0 {
		return 0
	}
	mid := len(inorder) / 2

	for i, j := 0, len(inorder)-1; i < mid && j > mid; i, j = i+1, j-1 {
		if inorder[i] != inorder[j] {
			return 0
		}
	}
	return 1
}

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
