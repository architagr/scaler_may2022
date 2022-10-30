package least_common_ancestor

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
 * @input C : Integer
 *
 * @Output Integer
 */
func Lca(A *treeNode, B int, C int) int {
	temp := A

	for temp != nil {
		if temp.value < B && temp.value < C {
			temp = temp.right
		} else if temp.value > B && temp.value > C {
			temp = temp.left
		} else {
			return temp.value
		}
	}
	return 0
}
