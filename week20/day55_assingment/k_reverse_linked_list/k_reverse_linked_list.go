package k_reverse_linked_list

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

func KReverseLinkedList(head *listNode, B int) *listNode {
	if B == 1 {
		return head
	}

	//Count length
	length := 0
	temp := head
	for temp != nil {
		temp = temp.next
		length++
	}

	//Initialize required states
	curr, next, prev := head, new(listNode), new(listNode)
	next, prev = nil, nil

	tempTail, tempHead, newHead := head, head, new(listNode)
	newHead = nil

	//Main logic
	for i := 0; i < length/B; i++ {
		for j := 0; j < B; j++ {
			next = curr.next
			curr.next = prev
			prev = curr
			curr = next
		}

		if i == 0 {
			newHead = prev
		} else {
			tempTail.next = prev
			tempTail = tempHead
		}

		tempHead = curr
		prev = nil
	}

	return newHead
}
