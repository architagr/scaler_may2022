package nearest_smaller_element

type Node struct {
	val  int
	prev *Node
}

func NewNode(value int) *Node {
	newNode := new(Node)
	newNode.val = value
	return newNode
}

type Stack struct {
	count int
	head  *Node
}

func (stack *Stack) Push(value int) {

	newNode := NewNode(value)
	stack.count++
	if stack.head == nil {
		stack.head = newNode
		return
	}
	newNode.prev = stack.head
	stack.head = newNode

}

func (stack *Stack) Pop() *Node {
	if stack.head == nil {
		return nil
	}
	node := stack.head
	stack.head = stack.head.prev
	stack.count--
	return node
}
func (stack *Stack) Top() int {
	if stack.head == nil {
		return -1
	}
	return stack.head.val
}
func prevSmaller(A []int) []int {
	result := make([]int, len(A))
	result[0] = -1
	stack := Stack{
		count: 0,
		head:  NewNode(A[0]),
	}
	for i := 1; i < len(A); i++ {
		if stack.Top() < A[i] {
			result[i] = stack.Top()
			stack.Push(A[i])
		} else {
			for stack.head != nil {
				if stack.Top() < A[i] {
					result[i] = stack.Top()
					stack.Push(A[i])
					break
				} else {
					stack.Pop()
				}
			}
			if stack.head == nil {
				result[i] = -1
				stack.Push(A[i])
			}
		}
	}
	return result
}
