package listcycle

type ListNode struct {
	value int
	next  *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	flag := false
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			flag = true
			break
		}
	}
	if flag {
		slow = head
		for fast != slow {
			fast = fast.next
			slow = slow.next
		}
		return fast
	}
	return nil
}
