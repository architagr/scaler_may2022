package sorted_array_to_balanced_bst

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
	next  *treeNode
}

func sortedArrayToBST(nums []int) *treeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	node := new(treeNode)
	node.value = nums[mid]
	node.left = sortedArrayToBST(nums[:mid])
	node.right = sortedArrayToBST(nums[mid+1:])
	return node
}
