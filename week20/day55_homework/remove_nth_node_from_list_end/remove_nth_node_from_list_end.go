package remove_nth_node_from_list_end

type listNode struct {
	value int
	next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.value = val
	node.next = nil
	return node
}

func RemoveNthFromEnd(A *listNode, B int) *listNode {
	count := countNodes(A)
	if B >= count {
		return A.next
	}
	temp := A
	for i := 1; i < count-B; i++ {
		temp = temp.next
	}
	temp.next = temp.next.next
	return A
}

func countNodes(head *listNode) int {
	count := 0
	if head == nil {
		return 0
	}

	for temp := head; temp != nil; temp = temp.next {
		count++
	}
	return count
}
