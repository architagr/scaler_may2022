package sum_of_min_and_max

type Node struct {
	data       int
	next, prev *Node
}

func NewNode(value int) *Node {
	node := new(Node)
	node.data = value
	return node
}

type DoulbllyLinkedList struct {
	count      int
	front, end *Node
}

func (ll *DoulbllyLinkedList) push(value int) {

	node := NewNode(value)

	if ll.count == 0 {
		ll.front = node
		ll.end = node
		ll.count++
		return
	}

	ll.end.next = node
	node.prev = ll.end
	ll.end = ll.end.next
	ll.count++
}
func (ll *DoulbllyLinkedList) popEnd() {
	if ll.count > 0 {
		ll.count--
		ll.end = ll.end.prev
		if ll.end != nil {
			ll.end.next = nil
		}
	}
	if ll.count == 0 {
		ll.end = nil
		ll.front = nil
	}
}

func (ll *DoulbllyLinkedList) popFront() {
	if ll.count > 0 {
		ll.count--
		ll.front = ll.front.next
		ll.front.prev = nil
	}
	if ll.count == 0 {
		ll.end = nil
		ll.front = nil
	}
}
func (ll *DoulbllyLinkedList) peekFront() int {
	return ll.front.data
}
func (ll *DoulbllyLinkedList) peekEnd() int {
	return ll.end.data
}
func SumMaxMin(A []int, B int) int {
	MOD := 1000000007
	min, max := 0, 0
	minStack := new(DoulbllyLinkedList)
	maxStack := new(DoulbllyLinkedList)
	ans := 0
	start := 0
	for i := 0; i < len(A); i++ {
		for minStack.count != 0 && minStack.peekEnd() > A[i] {
			minStack.popEnd()
		}
		minStack.push(A[i])

		for maxStack.count != 0 && maxStack.peekEnd() < A[i] {
			maxStack.popEnd()
		}
		maxStack.push(A[i])
		if i-start+1 >= B {

			min = (min + minStack.peekFront()) % MOD
			max = (max + maxStack.peekFront()) % MOD
			if A[start] == maxStack.peekFront() {
				maxStack.popFront()
			}
			if A[start] == minStack.peekFront() {
				minStack.popFront()
			}
			start++
		}

	}
	ans = (min + max) % MOD
	return (ans + MOD) % MOD

}
