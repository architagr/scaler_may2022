package sliding_window_maximum

type Node struct {
	data       int
	prev, next *Node
}

type DQueue struct {
	head, tail *Node
}

func (q *DQueue) Insert(data int) {
	node := new(Node)
	node.data = data
	if q.head == nil {
		q.head = node
		q.tail = node
	}
	node.prev = q.tail
	q.tail.next = node
	q.tail = node
}

func (q *DQueue) RemoveTop() *Node {
	if q.head == nil {
		return nil
	}

	node := q.head
	q.head.next.prev = nil
	q.head = q.head.next
	return node
}
func (q *DQueue) RemoveTail() *Node {
	if q.tail == nil {
		return nil
	}
	node := q.tail
	if q.tail == q.head {
		q.head = nil
		q.tail = nil

	} else {
		q.tail.prev.next = nil
		q.tail = q.tail.prev
	}
	return node
}

func (q *DQueue) Top() *Node {
	return q.head
}

func (q *DQueue) Tail() *Node {
	return q.tail
}

func slidingMaximum(A []int, B int) []int {
	q := new(DQueue)
	n := len(A)
	result := make([]int, 0, n-B+1)

	q.Insert(A[0])
	for i := 1; i < B; i++ {
		for node := q.Tail(); node != nil && node.data < A[i]; node = q.Tail() {
			q.RemoveTail()
		}
		q.Insert(A[i])
	}
	result = append(result, q.Top().data)
	for i := B; i < n; i++ {
		for node := q.Tail(); node != nil && node.data < A[i]; node = q.Tail() {
			q.RemoveTail()
		}
		q.Insert(A[i])
		if q.Top().data == A[i-B] {
			q.RemoveTop()
		}
		result = append(result, q.Top().data)

	}
	return result
}
