package counting_the_nodes

type treeNode struct {
	left, right *treeNode
	value       int
}

func CountNodes(A *treeNode) int {
	return count(A, 0)
}

func count(head *treeNode, max int) int {
	if head == nil {
		return 0
	}
	c := 0
	if head.value > max {
		c = 1
		max = head.value
	}
	return c + count(head.left, max) + count(head.right, max)
}
