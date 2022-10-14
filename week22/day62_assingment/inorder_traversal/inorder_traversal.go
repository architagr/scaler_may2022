package inorder_traversal

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
 * @Output Integer array.
 */
func InorderTraversal(A *treeNode) []int {
	return Parse(A)
}

func Parse(head *treeNode) []int {
	if head == nil {
		return []int{}
	}
	result := make([]int, 0)

	result = append(result, Parse(head.left)...)
	result = append(result, head.value)
	result = append(result, Parse(head.right)...)
	return result
}
