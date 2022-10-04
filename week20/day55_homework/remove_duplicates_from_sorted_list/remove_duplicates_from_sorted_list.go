package remove_duplicates_from_sorted_list

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
func DeleteDuplicates(A *listNode) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	for temp := A; temp.next != nil; {
		if temp.value == temp.next.value {
			temp.next = temp.next.next
		} else {
			temp = temp.next
		}
	}
	return A
}
