package middle_element_of_linked_list

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
 * @Output Integer
 */
func GetMiddleEle(head *listNode) int {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.next
	}
	temp = head

	mid := count / 2
	for i := 0; i < mid; i++ {
		temp = temp.next
	}
	return temp.value
}
