package odd_and_even_levels

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
var sumArr []int

func OddEvenDiff(A *treeNode) int {
	sumArr = make([]int, 2)
	sum(A, 0)
	return sumArr[0] - sumArr[1]
}
func sum(A *treeNode, level int) {
	if A == nil {
		return
	}
	sumArr[level%2] += A.value
	sum(A.left, level+1)
	sum(A.right, level+1)
}
