package sort_list

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

/**
 * @input A : Head pointer of linked list
 *
 * @Output head pointer of list.
 */
func SortList(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	mid := getMid(A)
	h2 := mid.next
	mid.next = nil
	A = SortList(A)
	h2 = SortList(h2)
	return mergeTwoLists(A, h2)
}

func getMid(head *listNode) *listNode {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	slow, fast := head, head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func mergeTwoLists(A *listNode, B *listNode) *listNode {
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
