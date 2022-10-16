package binary_tree_from_inorder_and_postorder

import "scaler-may-22/week22/day62_assingment/inorder_traversal"

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output root pointer of tree.
 */
var idx int = 0

func BuildTree(A []int, B []int) *inorder_traversal.TreeNode {
	idx = len(B) - 1
	return getTree(A, B, 0, len(B)-1)
}

func getTree(A, B []int, start, end int) *inorder_traversal.TreeNode {
	if start > end {
		return nil
	}
	curr := B[idx]
	root := inorder_traversal.TreeNode_new(curr)
	idx--
	if start != end {
		pos := search(A, start, end, curr)
		root.Right = getTree(A, B, pos+1, end)
		root.Left = getTree(A, B, start, pos-1)

	}
	return root
}

func search(A []int, start, end, cur int) int {
	for i := start; i <= end; i++ {
		if A[i] == cur {
			return i
		}
	}
	return -1
}
