package palindrome_list

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

func FindPalindrome(head *listNode) int {
	if head == nil {
		return 0
	}
	if head.next == nil {
		return 1
	}
	mid := getMid(head)

	list2 := reverseLinkedList(mid.next)
	temp := head
	for list2 != nil {
		if list2.value != temp.value {
			return 0
		}
		temp = temp.next
		list2 = list2.next
	}
	return 1
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
	temp.next = nil
	return p
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
