package reverse_linked_list

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

func ReverseLinkedList(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}
	p, c, n := head, head.next, head.next.next
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
