package merge_two_sorted_list

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

func MergeTwoLists(A *listNode, B *listNode) *listNode {
	if A == nil && B == nil {
		return nil
	} else if A == nil {
		return B
	} else if B == nil {
		return A
	}

	var h2 *listNode
	if A.value < B.value {
		h2 = A
		A = A.next
	} else {
		h2 = B
		B = B.next
	}

	temp := h2

	for A != nil && B != nil {
		if A.value < B.value {
			temp.next = A
			A = A.next
		} else {
			temp.next = B
			B = B.next
		}
		temp = temp.next
	}
	if A != nil {
		temp.next = A
	}
	if B != nil {
		temp.next = B
	}
	return h2
}
