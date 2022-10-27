package valid_binary_search_tree

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
 * @Output Integer
 */
func isValidBST(A *treeNode) int {
	arr := Parse(A)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return 0
		}
	}

	return 1
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
