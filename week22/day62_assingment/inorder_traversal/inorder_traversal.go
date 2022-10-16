package inorder_traversal

type TreeNode struct {
	Left  *TreeNode
	value int
	Right *TreeNode
}

func TreeNode_new(val int) *TreeNode {
	var node *TreeNode = new(TreeNode)
	node.value = val
	node.Left = nil
	node.Right = nil
	return node
}

/**
 * @input A : Root pointer of the tree
 *
 * @Output Integer array.
 */
func InorderTraversal(A *TreeNode) []int {
	return Parse(A)
}

func Parse(head *TreeNode) []int {
	if head == nil {
		return []int{}
	}
	result := make([]int, 0)

	result = append(result, Parse(head.Left)...)
	result = append(result, head.value)
	result = append(result, Parse(head.Right)...)
	return result
}
