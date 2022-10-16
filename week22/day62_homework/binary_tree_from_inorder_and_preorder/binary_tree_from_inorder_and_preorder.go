package binary_tree_from_inorder_and_preorder

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
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output root pointer of tree.
 */
var idx int = 0

func BuildTree(A []int, B []int) *treeNode {
	idx = 0
	return build(B, A, 0, len(A)-1)
}
func build(A, B []int, start, end int) *treeNode {
	if start > end || idx >= len(A) {
		return nil
	}
	curr := B[idx]
	idx++
	node := treeNode_new(curr)
	if start != end {
		pos := search(A, start, end, curr)
		node.left = build(A, B, start, pos-1)
		node.right = build(A, B, pos+1, end)
	}
	return node
}

func search(A []int, start, end, cur int) int {
	for i := start; i <= end; i++ {
		if A[i] == cur {
			return i
		}
	}
	return -1
}
