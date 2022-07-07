package design_linkelist

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

func Insert_node(head *listNode, position, value int) *listNode {
	// @params position, integer
	// @params value, integer
	newNode := new(listNode)
	newNode.value = value
	if head == nil {
		head = newNode
		return head
	}
	if position == 0 {
		newNode.next = head
		head = newNode
		return head
	}
	c := 0
	x := head
	for x.next != nil && c < position-1 {
		x = x.next
		c++
	}
	newNode.next = x.next
	x.next = newNode
	return head
}

func Delete_node(head *listNode, position int) *listNode {
	// @params position, integer
	if head == nil {
		return head
	}
	if position == 0 {
		head = head.next
		return head
	}
	c := 0
	x := head
	for x.next != nil && c < position-1 {
		x = x.next
		c++
	}

	if c <= position && x.next != nil {
		x.next = x.next.next
	}
	return head
}

const (
	InsertStart  = 0
	InsertEnd    = 1
	InsertMiddle = 2
	Delete       = 3
)

func solve(A [][]int) *listNode {
	var head *listNode
	length := 0
	for i := 0; i < len(A); i++ {
		switch A[i][0] {
		case InsertStart:
			head = Insert_node(head, 0, A[i][1])
			length++
		case InsertEnd:
			if head != nil {
				x := head
				for x.next != nil {
					x = x.next
				}
				x.next = listNode_new(A[i][1])
			} else {
				head = Insert_node(head, 0, A[i][1])
			}
			length++
		case InsertMiddle:
			if A[i][2] > length {
				continue
			} else {
				head = Insert_node(head, A[i][2], A[i][1])
				length++
			}
		case Delete:
			if length > 0 {
				head = Delete_node(head, A[i][1])
				length--
			}

		}
	}
	return head
}
