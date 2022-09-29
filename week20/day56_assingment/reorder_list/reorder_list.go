package reorder_list

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

func ReorderList(head *listNode) *listNode {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}

	mid := findMid(head)
	a1 := head
	a2 := mid.next
	mid.next = nil

	a2 = reverseLinkedList(a2)
	return mergeLinkedLists(a1, a2)
}

func mergeLinkedLists(a1, a2 *listNode) *listNode {
	t, t1, t2 := a1, a1.next, a2.next

	for a2 != nil {
		t.next = a2
		a2.next = t1
		t = t1
		a2 = t2
		if t1 != nil {
			t1 = t1.next
		}

		if t2 != nil {
			t2 = t2.next
		}
	}
	return a1
}

func findMid(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func reverseLinkedList(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}

	temp := head

	p, c, n := temp, temp.next, temp.next.next

	for c != nil {
		c.next = p
		p = c
		c = n
		if n != nil {
			n = n.next
		}
	}
	head.next = nil
	return p
}
