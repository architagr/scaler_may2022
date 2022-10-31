package kth_smallest_element_in_bst

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
func kthsmallest(A *treeNode, B int) int {
	arr := InOrderTrevarsal(A)
	return arr[B-1]
}

func InOrderTrevarsal(A *treeNode) []int {
	if A == nil {
		return []int{}
	}
	result := make([]int, 0)
	result = append(result, InOrderTrevarsal(A.left)...)
	result = append(result, A.value)
	result = append(result, InOrderTrevarsal(A.right)...)
	return result
}
