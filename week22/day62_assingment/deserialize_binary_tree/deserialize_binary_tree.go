package deserialize_binary_tree

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

func solve(A []int) *treeNode {
	n := len(A)
	queue := make([]*treeNode, 0, n)
	root := treeNode_new(A[0])
	x, i := 0, 1
	queue = append(queue, root)
	for i < n {
		n1, n2 := treeNode_new(A[i]), treeNode_new(A[i+1])
		i = i + 2
		if n1.value != -1 {
			queue[x].left = n1
			queue = append(queue, n1)
		}
		if n2.value != -1 {
			queue[x].right = n2
			queue = append(queue, n2)
		}
		x++

	}
	return root
}
