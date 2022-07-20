package bst_nodes_in_a_range

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func GetCount(A *treeNode, B int, C int) int {
	if A == nil {
		return 0
	}

	count := 0
	if A.value >= B && A.value <= C {
		count = 1
	}
	count += GetCount(A.left, B, C) + GetCount(A.right, B, C)

	return count
}
