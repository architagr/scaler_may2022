package path_to_given_node

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func FindPath(A *treeNode, B int) []int {
	path := make([]int, 0)
	calcPath(A, B, &path)
	return path
}

func calcPath(head *treeNode, B int, path *[]int) bool {
	if head == nil {
		return false
	}
	*path = append(*path, head.value)

	if head.value == B {
		return true
	}

	if calcPath(head.left, B, path) || calcPath(head.right, B, path) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
