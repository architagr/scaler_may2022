package perfect_number

type Node struct {
	data string
	next *Node
}

type Queue struct {
	head, tail *Node
}

func (q *Queue) Enqueue(data string) {
	node := new(Node)
	node.data = data

	if q.head == nil && q.tail == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}
func (q *Queue) Dequeue() (result *Node) {
	result = q.head
	if q.head == nil {
		return
	}
	q.head = q.head.next
	return
}
func PerfectNumber(A int) string {
	if A == 1 {
		return "11"
	}
	if A == 2 {
		return "22"
	}
	queue := new(Queue)
	queue.Enqueue("1")
	queue.Enqueue("2")

	ans := ""
	for i := 3; i <= A; i++ {
		a := queue.Dequeue().data
		ans = a + "1"
		queue.Enqueue(ans)

		if i == A {
			break
		}
		i++
		ans = a + "2"
		queue.Enqueue(ans)
	}

	return ans + string(reverse([]byte(ans)))
}

func reverse(a []byte) []byte {
	l, r := 0, len(a)-1

	for l <= r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return a
}
