package min_stack

type MinStack struct {
	head    *Node
	minHead *Node
}

type Node struct {
	data int
	next *Node
}

func NewNode(val int) *Node {
	n := new(Node)
	n.data = val
	return n
}

/** initialize your data structure here. */
func Constructor() *MinStack {
	data := new(MinStack)
	return data
}

func (this *MinStack) Push(val int) {

	n := NewNode(val)
	n.next = this.head
	this.head = n
	this.PushMin(val)
}

func (this *MinStack) PushMin(val int) {
	if this.minHead == nil {
		n := NewNode(val)
		this.minHead = n
	} else if this.minHead.data >= val {
		n := NewNode(val)
		n.next = this.minHead
		this.minHead = n

	}
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.PopMin(this.head.data)
	this.head = this.head.next
}

func (this *MinStack) PopMin(val int) {
	if this.minHead == nil {
		return
	}
	if this.minHead.data == val {
		this.minHead = this.minHead.next
	}
}

func (this *MinStack) Top() int {
	if this.head != nil {
		return this.head.data
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.minHead == nil {
		return -1
	}
	return this.minHead.data
}
