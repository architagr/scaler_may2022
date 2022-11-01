package common_nodes_in_two_bst

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
 * @input B : Root pointer of the tree
 *
 * @Output Integer
 */
var hashMap map[int]int = make(map[int]int)
var sum int = 0
var MOD = 1000000007

func CommonNodes(A *treeNode, B *treeNode) int {
	hashMap = make(map[int]int)
	sum = 0
	convertTreeToMap(A)
	check(B)
	return sum
}
func check(B *treeNode) {
	if B == nil {
		return
	}
	if x, ok := hashMap[B.value]; ok && x > 0 {
		hashMap[B.value]--
		sum = (sum + B.value) % MOD
	}
	check(B.right)
	check(B.left)
}
func convertTreeToMap(A *treeNode) {
	if A == nil {
		return
	}
	hashMap[A.value]++
	convertTreeToMap(A.left)
	convertTreeToMap(A.right)
}
