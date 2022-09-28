package delete_middle_node_of_a_linked_list

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
func solve(head *listNode) *listNode {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.next
	}
	temp = head
	if count == 1 {
		return nil
	}
	mid := count / 2
	for i := 1; i < mid; i++ {
		temp = temp.next
	}
	temp.next = temp.next.next
	return head
}
