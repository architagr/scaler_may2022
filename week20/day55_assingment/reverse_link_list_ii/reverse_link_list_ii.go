package reverse_link_list_ii

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
 * @input B : Integer
 * @input C : Integer
 *
 * @Output head pointer of list.
 */
func ReverseBetween(A *listNode, B int, C int) *listNode {
	if B == C {
		return A
	}

	dummy := listNode_new(-1)
	dummy.next = A
	temp := dummy

	for count := 0; count < B-1; count++ {
		temp = temp.next
	}
	C -= B

	t1 := temp.next
	p, c, n := t1, t1.next, t1.next.next
	for count := 0; count < C; count++ {
		c.next = p

		p = c
		c = n
		if n != nil {
			n = n.next
		}

	}
	t1.next = c
	temp.next = p

	return dummy.next

}
