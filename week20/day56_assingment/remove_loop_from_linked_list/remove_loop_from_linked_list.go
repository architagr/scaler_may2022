package remove_loop_from_linked_list

type ListNode struct {
	value int
	next  *ListNode
}

func RemoveLoop(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			break
		}
	}
	slow = head
	for fast.next != slow.next {
		fast = fast.next
		slow = slow.next
	}
	fast.next = nil
	return head
}
